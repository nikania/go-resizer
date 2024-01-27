import { ArrowUpIcon, DownloadIcon } from "@chakra-ui/icons";
import "./App.css";
import UploadFile from "./components/Upload";
import { Button, Stack } from "@chakra-ui/react";

function App() {
  const handleDownload = () => {
    const name: string = "img-3ecd19d4-0240-4bac-b6b7-c56c34547018.jpeg";
    fetch(`http://localhost:8080/download?name=${name}`)
      .then((resp) => {
        return resp.blob();
      })
      .then((blob) => {
        const link = document.createElement("a");
        link.download = name;
        link.href = URL.createObjectURL(blob);
        link.click();

        URL.revokeObjectURL(link.href);
      })
      .catch((err) => console.log(err));
  };
  const handleUpload = () => {
    let blob = new Blob(["Hello, world!"], { type: "text/plain" });
    fetch("http://localhost:8080/upload", {
      method: "POST",
      headers: {
        "Content-Type": "multipart/form-data;boundary=-",
      },
      body: blob,
    })
      .then((resp) => alert(resp.headers))
      .catch((err) => console.log(err));
  };
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
        <form
          method="post"
          // action="http://localhost:8080/upload"
          encType="multipart/form-data"
        >
          <div>
            <label htmlFor="file">Choose a file</label>
            <input type="file" id="file" name="myFile" />
          </div>
          <div>
            <Button
              type="submit"
              leftIcon={<ArrowUpIcon />}
              onClick={handleUpload}
            >
              Upload
            </Button>
          </div>
        </form>

        <Button onClick={handleResize}>Resize</Button>
        <Button onClick={handleConvert}>Convert</Button>
        <Button onClick={handleDelete}>Delete</Button>
      </Stack>
    </>
  );
}

export default App;
