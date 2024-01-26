import { ArrowUpIcon, DownloadIcon } from "@chakra-ui/icons";
import "./App.css";
import UploadFile from "./components/Upload";
import { Button, Stack } from "@chakra-ui/react";

function App() {
  const handleDownload = () => {
    fetch(
      "http://localhost:8080/download?name=img-3ecd19d4-0240-4bac-b6b7-c56c34547018.jpeg",
    )
      .then((resp) => resp.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
  };
  const handleUpload = () => {};
  function handleResize(
    event: MouseEvent<HTMLButtonElement, MouseEvent>,
  ): void {
    fetch(
      "http://localhost:8080/resize?name=img-3ecd19d4-0240-4bac-b6b7-c56c34547018.jpeg&width=100&height=500",
    )
      .then((resp) => resp.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
  }

  function handleConvert(
    event: MouseEvent<HTMLButtonElement, MouseEvent>,
  ): void {
    fetch(
      "http://localhost:8080/convert?name=img-3ecd19d4-0240-4bac-b6b7-c56c34547018.jpeg&to=image/gif",
    )
      .then((resp) => resp.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
  }

  function handleDelete(
    event: MouseEvent<HTMLButtonElement, MouseEvent>,
  ): void {
    fetch(
      "http://localhost:8080/delete?name=img-3ecd19d4-0240-4bac-b6b7-c56c34547018.jpeg",
    )
      .then((resp) => resp.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
  }

  return (
    <>
      <h1>Resizing app</h1>
      <p>todo</p>
      <p>axios or any other async backend calls</p>
      <p>components structure</p>
      <p>page layout</p>

      <br></br>
      <UploadFile />
      <Stack direction="row" spacing={4}>
        <Button
          leftIcon={<DownloadIcon />}
          onClick={handleDownload}
          colorScheme="teal"
          variant="solid"
        >
          Download
        </Button>
        <Button leftIcon={<ArrowUpIcon />} onClick={handleUpload}>
          Upload
        </Button>
        <Button onClick={handleResize}>Resize</Button>
        <Button onClick={handleConvert}>Convert</Button>
        <Button onClick={handleDelete}>Delete</Button>
      </Stack>
    </>
  );
}

export default App;
