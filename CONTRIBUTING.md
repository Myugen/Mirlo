# CONTRIBUTING

## How to contribute

If you don't know where to begin, take a look at the [project issues](https://github.com/alephshahor/Mirlo/issues), there you'll find a list of tasks yet to be done. If you wish to propose a feature that's not in the list please open a pull request introducing your proposal and wait for it to be addressed. 

Once you've finished the task implementation open a pull request (if it doesn't exist yet).

The following conditions must be required:

* Only modify the files related with the task that you've been assigned to fix.
* Be clear and concise with the code style.
* Write test for whatever new functionality that you add / modify.
* Make sure that the tests pass.
* Add a summary of what you've done to complete the task

## Commit guidelines

In Mirlo we use the [Angular Commit Message Format](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#commit). The scope we're specifying is the most nested folder in which the file is located, so for example if we have a directory structure like:

```
+ server
+-- subfolder1
	+-- subfolder2
		+-- target_file
```

The commit message will have a scope called `subfolder2`.  

**IMPORTANT**: Those functionalities that affect all the app use the _all_ scope.

If you want to add more information about what you've done (surpass the 100 characters) do it in the following way:

```
feat(some_scope): This is a summary
> Did something
> Did another thing
> ...
```





