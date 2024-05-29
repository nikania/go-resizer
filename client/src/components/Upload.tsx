import { ArrowUpIcon } from "@chakra-ui/icons";
import { Button, Stack } from "@chakra-ui/react";
import { useContext } from "react";
import { FileContextType, FileContext } from "../context/FileContext";

// react component for uploading files on server
const UploadFile = () => {
  const { name, changeName } = useContext<FileContextType>(FileContext);
  console.log("🚀 ~ UploadFile ~ name:", name);
  const handleFileUpload = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const fileInput = document.getElementById("file") as HTMLInputElement;
    const formData = new FormData();
    formData.append("file", fileInput.files[0]);

    fetch("http://localhost:8080/upload", {
      method: "POST",
      body: formData,
    })
      .then((resp) => {
        console.log(resp);
        alert("File uploaded successfully");
        return resp.json();
      })
      .then((data: { name: string }) => {
        console.log("🚀 ~ .then ~ data:", data);
        changeName(data.name);
      })
      .catch((err) => console.log(err));
  };
  console.log("🚀 ~ UploadFile ~ name:", name);

  return (
    <>
      <Stack direction="row" spacing={4}>
        <form
          method="post"
          // action="http://localhost:8080/upload"
          onSubmit={handleFileUpload}
          encType="multipart/form-data"
        >
          <div>
            <label htmlFor="file">Choose a file</label>
            <input type="file" id="file" name="file" />
          </div>
          <div>
            <Button type="submit" leftIcon={<ArrowUpIcon />}>
              Upload
            </Button>
          </div>
        </form>
      </Stack>
    </>
  );
};

export default UploadFile;
