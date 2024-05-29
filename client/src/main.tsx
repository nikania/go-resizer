import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { ChakraProvider } from "@chakra-ui/react";
import { FileProvider } from "./context/FileContext.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ChakraProvider>
      <FileProvider>
        <App />
      </FileProvider>
    </ChakraProvider>
  </React.StrictMode>,
);
