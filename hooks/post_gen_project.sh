#!/bin/bash
echo "Running post_gen_project.sh"

# Check if the folder exists
if [ -d ".github" ]; then
    if [ ! -d "../.github/workflows" ]; then
        echo "Creating folder ../.github/workflows"
        mkdir -p ../.github/workflows
        cp .github/workflows/* ../.github/workflows
    else
        echo "Folder ../.github/workflows already exists"
    fi

    rm -rf .github
    echo "GitHub Workflows moved successfully"
else
    echo "Folder not found!"
fi

if [ ! -f "../Makefile" ]; then
    echo "Creating root Makefile"
    mv rootMakefile ../Makefile
    echo -en "\n" >> ../Makefile
    cat Makefile >> ../Makefile
    rm Makefile
else
    echo "root Makefile already exists"
    echo -en "\n" >> ../Makefile
    cat Makefile >> ../Makefile
    rm rootMakefile
    rm Makefile
fi

if [ ! -f "../go.mod" ]; then
    echo "Creating go.mod"
    mv go.mod ../
else
    echo "Makefile already exists"
    rm go.mod
fi

if [ ! -f "../template.yaml" ]; then
    echo "Creating rootTemplate.yaml"
    mv rootTemplate.yaml ../template.yaml
else
    echo "rootTemplate.yaml already exists"
    rm rootTemplate.yaml

    yaml_file="../template.yaml"

    # Specify the item you want to check
    projectResourceKey="{{cookiecutter.module_name | replace('-', '_')}}"
    newYamlResourceKey="Resources.$projectResourceKey"

    # Check if the item exists in the YAML file
    item_exists=$(yq ".Resources | has(\"$projectResourceKey\")" $yaml_file)

    if [ "$item_exists" == "true" ]; then
        echo "Item $newYamlResourceKey already exists in the root template file. Skipping... PLEASE ADD IT MANUALLY!"
    else
        temp_file=$(mktemp)
        yaml_object_to_add=$(cat <<EOF
Type: AWS::Serverless::Application
Properties:
  Location: ./{{cookiecutter.module_name}}/template.yaml
  # Optional parameter that can have default value overridden
  Parameters:
    Project: !Ref Project
    ProjectGroup: !Ref ProjectGroup
    Environment: !Ref Environment
    TracingName: !Ref TracingName
EOF
)
        # Add the item and its value to the YAML file
        echo "$yaml_object_to_add" > "$temp_file"
        yq eval-all "select(fileIndex==0).$newYamlResourceKey = select(fileIndex==1) | select(fileIndex==0)" $yaml_file $temp_file -i $yaml_file

        echo "Item added successfully!"
    fi

fi

if [ ! -f "../.gitignore" ]; then
    echo "Creating root .gitIgnore"
    mv rootGitIgnore ../.gitignore
else
    rm rootGitIgnore
fi