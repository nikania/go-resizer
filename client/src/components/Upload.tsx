import { ArrowUpIcon } from "@chakra-ui/icons";
import { Button, Stack } from "@chakra-ui/react";

// react component for uploading files on server
const UploadFile = () => {
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

  return (
    <>
      <Stack direction="row" spacing={4}>
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
      </Stack>
    </>
  );
};

export default UploadFile;
