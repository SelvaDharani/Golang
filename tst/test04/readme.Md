**_ usage of following methods to signal a failure _**

Error : will log and marks the test function as failed. Execution will continue.

Errorf : will log (with the specified format) and marks the test function as failed. Execution will continue.

Fail : will mark the function as failing. Execution will continue.

FailNow : this one will mark the test as failed and stop the execution of the current test function (if you have other assertions, they will not be tested).
