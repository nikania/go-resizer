import { Button } from "@chakra-ui/react";

export default function Delete() {
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
      <Button onClick={handleDelete}>Delete</Button>
    </>
  );
}
