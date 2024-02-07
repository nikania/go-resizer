import { Button, Heading } from "@chakra-ui/react";
import PageLayout from "../layouts/PageLayout";
import { useState } from "react";

export default function Resize() {
  const [name, setName] = useState("");
  const [width, setWidth] = useState(100);
  const [height, setHeight] = useState(100);

  console.log(name, width, height);

  function handleResize(event): void {
    event.preventDefault();
    fetch(
      `http://localhost:8080/resize?name=${name}&width=${width}&height=${height}`,
    )
      .then((resp) => resp.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
  }
  return (
    <PageLayout>
      <Heading>Resize image</Heading>
      <form onSubmit={handleResize}>
        <label>
          <span>FileName:</span>
          <input
            type="text"
            onChange={(e) => setName(e.target.value)}
            value={name}
          />
        </label>
        <label>
          <span>width:</span>
          <input
            type="number"
            onChange={(e) => setWidth(Number(e.target.value))}
            value={width}
          />
        </label>
        <label>
          <span>height:</span>
          <input
            type="number"
            onChange={(e) => setHeight(Number(e.target.value))}
            value={height}
          />
        </label>
        <Button type="submit">Resize</Button>
      </form>
      {/* <Button onClick={handleResize}>Resize</Button> */}
    </PageLayout>
  );
}
