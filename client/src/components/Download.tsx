import { DownloadIcon } from "@chakra-ui/icons";
import { Button } from "@chakra-ui/react";

const DownloadFile = () => {
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

  return (
    <Button
      leftIcon={<DownloadIcon />}
      onClick={handleDownload}
      colorScheme="teal"
      variant="solid"
    >
      Download
    </Button>
  );
};

export default DownloadFile;
