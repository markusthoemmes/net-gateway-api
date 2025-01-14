#!/usr/bin/env bash

# Copyright 2021 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

CONTOUR_OPERATOR_VERSION="v1.19.0" # This is for controlling which version of contour-operator we want to use.
YAML_URL="https://raw.githubusercontent.com/projectcontour/contour-operator/${CONTOUR_OPERATOR_VERSION}/examples/operator/operator.yaml"

kubectl apply -f $YAML_URL
