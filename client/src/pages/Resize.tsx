import { Button, Heading } from "@chakra-ui/react";
import DownloadFile from "../components/Download";
import PageLayout from "../layouts/PageLayout";

export default function Resize() {
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
  return (
    <PageLayout>
      <Heading>Resize image</Heading>
      <Button onClick={handleResize}>Resize</Button>
    </PageLayout>
  );
}
