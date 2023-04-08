/* Copyright 2023 Vítor de Albuquerque Torreão

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * 	http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/fx"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	return fmt.Sprintf("Hello, %s", event.Name), nil
}

type LambdaService struct{}

func NewLambdaService(lc fx.Lifecycle) *LambdaService {
	srv := LambdaService{}
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go lambda.Start(HandleRequest)
			return nil
		},
	})
	return &srv
}

func main() {
	fx.New(
		fx.Provide(NewLambdaService),
		fx.Invoke(func(*LambdaService) {}),
	).Run()
}
