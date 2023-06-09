# Copyright 2023 Vítor de Albuquerque Torreão

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.20.2-alpine3.16 as build

WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN go build

# Copy artifacts to a clean image

FROM alpine:3.16
COPY --from=build /app/crcler-go /crcler-go
ENTRYPOINT [ "/crcler-go" ]
