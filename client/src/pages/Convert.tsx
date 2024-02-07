import { Button, Heading, Input } from "@chakra-ui/react";
import PageLayout from "../layouts/PageLayout";
import { FormControl, FormLabel, Select } from "@chakra-ui/react";
import { useState } from "react";

export default function Convert() {
  const [name, setName] = useState("");
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
    <PageLayout>
      <Heading>Convert image</Heading>
      <FormControl>
        <FormLabel>
          <span>FileName:</span>
          <Input
            type="text"
            onChange={(e) => setName(e.target.value)}
            value={name}
          />
        </FormLabel>
        <FormLabel>Convert image</FormLabel>
        <Select
          placeholder="Convert To"
          onChange={(e) => setTo(e.target.value)}
        >
          <option value="image/jpg">JPG</option>
          <option value="image/png">PNG</option>
          <option value="image/gif">GIF</option>
        </Select>
      </FormControl>
      <Button onClick={handleConvert}>Convert</Button>
    </PageLayout>
  );
}
