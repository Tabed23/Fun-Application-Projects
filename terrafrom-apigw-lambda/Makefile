.PHONY: build clean


build: lambda_function/lambda_function.py
	zip -r9 lambda.zip lambda_function/lambda_function.py 

clean:
	rm -f lambda_function.zip


init:
	terraform init

plan:
	terraform plan -var-file=variable.tfvars


apply:
	terraform apply -var-file=variable.tfvars --auto-approve

dis:
	terraform destroy -var-file=variable.tfvars --auto-approve
