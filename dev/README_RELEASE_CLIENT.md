<!--
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 -->

# Release Process

Typically, releases are done coinciding with major and minor releases to Airflow. Therefore, a release of (for example) 
2.3.0 of this client would correspond with 2.3.X of Airflow.

The go client is generated using Airflow's [openapi spec](https://github.com/apache/airflow/blob/master/clients/gen/go.sh).
To update the client for new APIs do the following steps:

```bash
# clone this repo
git clone git@github.com:apache/airflow-client-go.git
cd airflow-client-go
export CLIENT_REPO_ROOT=$(pwd)
cd ..

# clone Airflow repo (if not already)
git clone git@github.com:apache/airflow.git
cd airflow
export AIRFLOW_REPO_ROOT=$(pwd)
```
Edit the file `airflow/airflow/api_connexion/openapi/v1.yaml`
Make sure it has the following `securitySchema`s listed under security `section`
```yaml
security:
  - Basic: []
  - GoogleOpenId: []
  - Kerberos: []
```
If your deployment of Airflow uses any different authentication mechanism than the three listed above, you might need to make further changes to the `v1.yaml` and generate your own client, see [OpenAPI Schema specification](https://swagger.io/docs/specification/authentication/) for details.
(*These changes should not be commited to the upstream `v1.yaml` [as it will generate misleading openapi documentation](https://github.com/apache/airflow/pull/17174)*)


```bash

# bump up the version in go.sh & run the following command
${AIRFLOW_REPO_ROOT}/clients/gen/go.sh airflow/api_connexion/openapi/v1.yaml ${CLIENT_REPO_ROOT}/airflow

cd ${CLIENT_REPO_ROOT}
```

- Get a diff between the last airflow version and the current airflow version
  which this release is based on:

      ```shell script
      git log 2.3.0..2.5.0 --pretty=oneline
      ```

- Update CHANGELOG.md with the details.

- Raise a PR in airflow-client-go. **Don't merge the PR as this would act as a prod release**.

# Prepare the Apache Airflow Go RC

## RC Sources

Go client is released as source. The generated source we vote upon should be the exact ones we vote against, without any modification than
renaming – i.e. the contents of the files must be the same between voted release candidate and final release.

- Checkout the branch of your PR. (with the newly generated source for the Go client)

- Set environment variables

    ```shell script
    # Set Version
    export VERSION=2.0.0rc1
    export VERSION_WITHOUT_RC=${VERSION%rc?}
    # Set the airflow version that this release is based
    export AIRFLOW_VERSION=2.1.4


    # Example after cloning
    git clone https://github.com/apache/airflow-client-go.git
    cd airflow-client-go
    export CLIENT_REPO_ROOT=$(pwd)
    ```

- Tag your RC and push it

    ```shell script
    git tag -s ${VERSION}
    git push origin ${VERSION}
    ```

- Clean the checkout

    ```shell script
    git clean -fxd
    ```

- Tarball the repo

    ```shell script
    git archive --format=tar.gz ${VERSION} --prefix=apache-airflow-client-${VERSION_WITHOUT_RC}/ -o apache-airflow-client-${VERSION_WITHOUT_RC}-source.tar.gz
    mkdir dist
    mv apache-airflow-client-${VERSION_WITHOUT_RC}-source.tar.gz dist
    ```

- Generate SHA512/ASC (If you have not generated a key yet, generate it by following instructions on
  http://www.apache.org/dev/openpgp.html#key-gen-generate-key)

    ```shell script
    pushd dist
    ${CLIENT_REPO_ROOT}/dev/sign.sh *
    popd
    ```

- Push the artifacts to ASF dev dist repo

```shell script
# First clone the repo
svn checkout https://dist.apache.org/repos/dist/dev/airflow airflow-dev

# Create new folder for the release
cd airflow-dev/clients/go
svn mkdir ${VERSION}

# Move the artifacts to svn folder & commit
mv ${CLIENT_REPO_ROOT}/dist/apache{-,_}*client-${VERSION_WITHOUT_RC}* ${VERSION}/
cd ${VERSION}
svn add *
svn commit -m "Add artifacts for Apache Airflow Go Client ${VERSION}"
cd ${CLIENT_REPO_ROOT}
rm -rf airflow-dev
```

## Prepare Vote email on the Airflow Client release candidate

Subject:

```shell script
cat <<EOF
[VOTE] Release Airflow Go Client ${VERSION_WITHOUT_RC} from ${VERSION}
EOF
```

Body:

```shell script
cat <<EOF
Hey fellow Airflowers,

I have cut the first release candidate for the Airflow Go Client ${VERSION}.
The client consists of APIs corresponding to REST APIs available in
*Apache Airflow ${AIRFLOW_VERSION}*. This email is calling for a vote on
the release, which will last for 72 hours. Consider this my (binding) +1.

Airflow Client ${VERSION} is available at:
https://dist.apache.org/repos/dist/dev/airflow/clients/go/${VERSION}/

Or also available on github:
https://github.com/apache/airflow-client-go/tree/${VERSION}/

*apache-airflow-client-${VERSION}-source.tar.gz* is a source release that comes with
INSTALL instructions.

Public keys are available at:
https://dist.apache.org/repos/dist/release/airflow/KEYS

Only votes from PMC members are binding, but the release manager should
encourage members of the community to test the release and vote with
"(non-binding)".

*Changelog:*

*Major changes:*
...

*Major fixes:*
...

*New API supported:*
...

Cheers,
<your name>
EOF
```

# Verify the release candidate by PMCs
See Airflow process documented [here](https://github.com/apache/airflow/blob/master/dev/README_RELEASE_AIRFLOW.md#verify-the-release-candidate-by-pmcs).

## SVN check
See Airflow process documented [here](https://github.com/apache/airflow/blob/master/dev/README_RELEASE_AIRFLOW.md#svn-check) (just replace Airflow with Airflow Client).

## Licence check
See Airflow process documented [here](https://github.com/apache/airflow/blob/master/dev/README_RELEASE_AIRFLOW.md#licence-check).

## Signature check
See Airflow process documented [here](https://github.com/apache/airflow/blob/master/dev/README_RELEASE_AIRFLOW.md#signature-check).

# Verify release candidates by Contributors
This can be done (and we encourage to) by any of the Contributors. In fact, it's best if the
actual users of Airflow Client test it in their own staging/test installations. Each release candidate
is available on github apart from SVN packages, so everyone should be able to install
the release candidate version of Airflow Client via simply (<VERSION> is 2.0.0 for example, and <X> is
release candidate number 1,2,3,....).

```shell script
go get -v  github.com/apache/airflow-client-go@<VERSION>rc<X>
```

Once you install and run Airflow Client, you should perform any verification you see as necessary to check
that the client works as you expected.

# Publish the final Apache Airflow client release

## Summarize the voting for the Apache Airflow client release

```shell script
Hello,

Apache Airflow go Client 2.5.0 (based on RC1) has been accepted.

3 "+1" binding votes received:
- Ephraim Anierobi
- Jarek Potiuk
- Jed Cunningham


1 "+1" non-binding votes received:

- Pierre Jeambrun

Vote thread:
https://lists.apache.org/thread/1qcj0r67dff3zg0w2vyfhr30fx9xtp3y

I'll continue with the release process, and the release announcement will follow shortly.

Cheers,
<your name>
```

## Publish release to SVN

```shell script
# Go to Airflow go client sources first
cd <YOUR_AIRFLOW_CLIENT_REPO_ROOT>
export CLIENT_REPO_ROOT="$(pwd)"
cd ..
# Clone the AS
[ -d asf-dist ] || svn checkout --depth=immediates https://dist.apache.org/repos/dist asf-dist
svn update --set-depth=infinity asf-dist/{release,dev}/airflow
CLIENT_DEV_SVN="${PWD}/asf-dist/dev/airflow/clients/go"
CLIENT_RELEASE_SVN="${PWD}/asf-dist/release/airflow/clients/go"
cd "${CLIENT_RELEASE_SVN}"

export RC=2.0.0rc1
export VERSION=${RC/rc?/}

# Create new folder for the release
svn mkdir ${VERSION}
cd ${VERSION}

# Move the artifacts to svn folder & commit
for f in ${CLIENT_DEV_SVN}/$RC/*; do 
  svn cp $f . ; 
done
svn commit -m "Release Apache Airflow Go Client ${VERSION} from ${RC}"

# Remove old release
# http://www.apache.org/legal/release-policy.html#when-to-archive
cd ..
export PREVIOUS_VERSION=1.0.0
svn rm ${PREVIOUS_VERSION}
svn commit -m "Remove old release: ${PREVIOUS_VERSION}"
```

Verify that the packages appear in [airflow](https://dist.apache.org/repos/dist/release/airflow/clients/go)

- Checkout your PR

- Push Tag for the final version

    ```shell script
    git tag ${VERSION}
    git push origin ${VERSION}
    ```

- Merge the PR

## Create release on GitHub

Create a new release on GitHub with the release notes and assets from the release svn.

## Notify developers of release

See Airflow process documented [here](https://github.com/apache/airflow/blob/master/dev/README_RELEASE_AIRFLOW.md#notify-developers-of-release) (just replace Airflow with Airflow Client)

## Add release data to Apache Committee Report Helper

Add the release data (version and date) at: https://reporter.apache.org/addrelease.html?airflow
