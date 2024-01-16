import "./App.css";
import UploadFile from "./components/Upload";
import { Button, Upload } from "antd";

function App() {
  return (
    <>
      <h1>Resizing app</h1>
      <p>todo</p>
      <Button type="primary">Button</Button>
      <Upload />
      <br></br>
      <UploadFile />
    </>
  );
}

export default App;
