import { Button, Heading } from "@chakra-ui/react";
import PageLayout from "../layouts/PageLayout";

export default function Convert() {
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
  return (
    <PageLayout>
      <Heading>Convert image</Heading>

      <Button onClick={handleConvert}>Convert</Button>
    </PageLayout>
  );
}
