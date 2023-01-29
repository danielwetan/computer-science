<?php
interface ActionHandler
{
    public function execute();
}

#[Attribute]
class SetUp {}

#[SecondAttribute]
class Something {}

class CopyFile implements ActionHandler
{
    public string $fileName;
    public string $targetDirectory;

    #[SetUp]
    public function testing() {
      echo "it works";
    }

    #[SetUp]
    public function another() {
      echo "awesome";
    }

    #[SetUp]
    #[Something]
    public function fileExists()
    {
        if (!file_exists($this->fileName)) {
            throw new RuntimeException("File does not exist");
        }
    }

    #[SetUp]
    public function targetDirectoryExists()
    {
        if (!file_exists($this->targetDirectory)) {
            mkdir($this->targetDirectory);
        } elseif (!is_dir($this->targetDirectory)) {
            throw new RuntimeException("Target directory $this->targetDirectory is not a directory");
        }
    }

    public function execute()
    {
        copy($this->fileName, $this->targetDirectory . '/' . basename($this->fileName));
    }
}

function executeAction(ActionHandler $actionHandler)
{
    $reflection = new ReflectionObject($actionHandler);

    foreach ($reflection->getMethods() as $method) {
      
      $attributes = $method->getAttributes(SetUp::class);      
      $secondAttributes = $method->getAttributes(Something::class);
      
      print_r($attributes);
      print_r($secondAttributes);

        if (count($attributes) > 0) {
            $methodName = $method->getName();


            $actionHandler->$methodName();
        }
    }

    $actionHandler->execute();
}


$copyAction = new CopyFile();
$copyAction->fileName = "/home/daniel/code/computer-science/README.md";
$copyAction->targetDirectory = "/home/daniel/code/computer-science/temp";

executeAction($copyAction);