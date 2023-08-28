build:
	sam build --beta-features

deploy: build
	sam deploy \
		--stack-name={{cookiecutter.stack_name}} \
		--no-confirm-changeset --no-fail-on-empty-changeset \
		--resolve-s3 \
		--capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM CAPABILITY_AUTO_EXPAND