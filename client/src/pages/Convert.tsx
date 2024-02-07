import { Button, Heading, Input } from "@chakra-ui/react";
import PageLayout from "../layouts/PageLayout";
import { FormControl, FormLabel, Select } from "@chakra-ui/react";
import { useState, useContext } from "react";
import { FileContext, FileProvider } from "../context/FileContext";

export default function Convert() {
  const { name } = useContext(FileContext);
  const [to, setTo] = useState("");

  function handleConvert(
    event: MouseEvent<HTMLButtonElement, MouseEvent>,
  ): void {
    event.preventDefault();
    fetch(`http://localhost:8080/convert?name=${name}&to=${to}`)
      .then((resp) => resp.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
  }
  return (
    <FileProvider>
      <PageLayout>
        <Heading>Convert image</Heading>
        <FormControl>
          <FormLabel>Convert image</FormLabel>
          <Select
            placeholder="Convert To"
            onChange={(e) => setTo(e.target.value)}
          >
            <option value="image/jpeg">JPG</option>
            <option value="image/png">PNG</option>
            <option value="image/gif">GIF</option>
          </Select>
        </FormControl>
        <Button onClick={handleConvert}>Convert</Button>
      </PageLayout>
    </FileProvider>
  );
}
