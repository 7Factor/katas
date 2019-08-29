# Java Katas

## Compile

`$ javac {ClassFile.java} && javac -cp junit-4.12.jar:hamcrest-core-1.3.jar:. {TestClassFile.java}`

For example, `java -cp junit-4.12.jar:hamcrest-core-1.3.jar:. org.junit.runner.JUnitCore WrapperTest`

## Running Tests

`$ java -cp junit-4.12.jar:hamcrest-core-1.3.jar:. org.junit.runner.JUnitCore {TestClass}`

For example, `java -cp junit-4.12.jar:hamcrest-core-1.3.jar:. org.junit.runner.JUnitCore WrapperTest` to run the tests for the wordwrap kata.