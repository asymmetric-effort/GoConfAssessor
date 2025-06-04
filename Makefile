.PHONY: tree
tree:
	@echo 'create build/file_structure.txt'
	@treehash > build/file_structure.txt
	#@tree > build/file_structure.txt

.PHONY: zip
zip: tree
	@echo 'create build/project_files.zip'
	zip -r ./build/project.zip . -x ".git/*" "build/*"

clean:
	@rm -rf ./build &> /dev/null || true
	@mkdir -p ./build &> /dev/null || true