import { Button, Heading } from "@chakra-ui/react";
import PageLayout from "../layouts/PageLayout";
import { FormControl, FormLabel, Select } from "@chakra-ui/react";
import { useState, useContext, MouseEvent } from "react";
import { FileContext, FileContextType } from "../context/FileContext";

export default function Convert() {
  const { name } = useContext<FileContextType>(FileContext);
  const [to, setTo] = useState("");

  function handleConvert(
    event: MouseEvent<HTMLButtonElement, MouseEvent>,
  ): void {
    event.preventDefault();
    fetch(`http://localhost:8080/convert?name=${name}&to=${to}`)
      .then((resp) => {
        const contentDisposition =
          resp.headers.get("Content-Disposition") || "";
        const filename =
          contentDisposition.split("filename=")[1] || "defaultFilename.txt";

        return resp.blob().then((blob) => ({ blob, filename }));
      })
      .then(({ blob, filename }) => {
        const link = document.createElement("a");
        link.download = filename;
        link.href = URL.createObjectURL(blob);
        link.click();

        URL.revokeObjectURL(link.href);
      })
      .catch((err) => console.log(err));
  }
  return (
    <PageLayout>
      <Heading>Convert image</Heading>
      <form onSubmit={handleConvert}>
        <FormLabel>Convert image</FormLabel>
        <Select
          placeholder="Convert To"
          onChange={(e) => setTo(e.target.value)}
        >
          <option value="image/jpeg">JPG</option>
          <option value="image/png">PNG</option>
          <option value="image/gif">GIF</option>
        </Select>
        <Button type="submit">Convert</Button>
      </form>
    </PageLayout>
  );
}
