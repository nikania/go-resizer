import { Button } from "antd";
import React from "react";

// react component for uploading files on server
const UploadFile = () => {
  function handleUpload(event: MouseEvent<HTMLElement, MouseEvent>): void {
    throw new Error("Function not implemented.");
  }

  return (
    <Button onClick={handleUpload} type="primary">
      Upload file
    </Button>
  );
};

export default UploadFile;
