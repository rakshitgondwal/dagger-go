# Buildsafe for CI/CD tools

In this tutorial, we're going to learn how to use Buildsafe with your CI/CD tools.
For this tutorial, we're going to be demoing Buildsafe with Dagger, to lint our application.

## Pre-requisites

- [Docker](https://docs.docker.com/engine/install/)
- [Buildsafe](https://github.com/buildsafedev/bsfrelease/blob/main/docs/quickstart.md#setup)
- [Dagger](https://docs.dagger.io/install)

## Follow along:

### Setting up a basic application

We're going to be using a simple Golang CLI application for this demo.
You can use the below code to setup the application.

```go
func main() {
	var greeting string

	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			if greeting != "" {
				fmt.Println(greeting)
			} else {
				fmt.Println(viper.GetString("greeting"))
			}
		},
	}

	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting message")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Setting up Buildsafe to build our image

I assume that you have bsf already installed, if not, please check out the link in the pre-requisites.

- Run the following command to setup the initial application.
  ```    
   bsf init
  ```
  This is will setup all the initial files in your root directory for bsf to work.

- For this tutorial, we're going to be focussing on the export part of bsf.
  If you'd like to know how other functionalities of bsf works, you can check out this link for a
  [quickstart](https://github.com/buildsafedev/bsfrelease/blob/main/docs/quickstart.md) guide.

- For our usecase, we'll be requiring a linter tool to lint our application during CI process, so we'll
  need to add a `golangci-lint` pacakge inside our development pacakges in `bsf.hcl`.
  You can add it via running the following command:
  ```    
   bsf search golangci-lint
  ```

- For exporting an image for our application, we're going to require an export block in our `bsf.hcl` file.
  You can go ahead and add this code block to your `bsf.hcl` file.
  ```    
   export "dev" {
    artifactType = "oci"
    name         = "golang-tut/greeting-cli"
    cmd          = ["/result/bin/hello", "--greeting","Good morning!"]
    platform     = "linux/amd64"
    devDeps      = true
    config       = "."
   }
  ```

- The thing that requires our attention over here is the `devDeps` and `config` fields in the export block.

  `devDeps` if set to true will export all our development packages in our OCI image. By default we only copy
  over the runtime packages in our OCI image.

  `config` field is defined as the path to the directory of your application that you want to copy over to the image.

- To build the image run the following command:
  ```
   bsf export dev
  ```
