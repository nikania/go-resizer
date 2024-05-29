import { Button, Heading } from "@chakra-ui/react";
import PageLayout from "../layouts/PageLayout";
import { useState, useContext } from "react";
import {
  FileContext,
  FileContextType,
  // FileProvider,
} from "../context/FileContext";

export default function Resize() {
  const { name } = useContext<FileContextType>(FileContext);
  console.log("🚀 ~ Resize ~ name:", name);
  const [width, setWidth] = useState(100);
  console.log("🚀 ~ Resize ~ width:", width);
  const [height, setHeight] = useState(100);

  function handleResize(event: { preventDefault: () => void }): void {
    console.log("🚀 ~ Resize ~ name:", name);
    event.preventDefault();
    fetch(
      `http://localhost:8080/resize?name=${name}&width=${width}&height=${height}`,
    )
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
  }
  return (
    <PageLayout>
      <Heading>Resize image</Heading>
      <form onSubmit={handleResize}>
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
