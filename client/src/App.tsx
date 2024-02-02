import "./App.css";
import {
  createBrowserRouter,
  RouterProvider,
  createRoutesFromElements,
  Route,
} from "react-router-dom";
import RootLayout from "./layouts/RootLayout.tsx";
import Home from "./pages/Home.tsx";
import Resize from "./pages/Resize.tsx";
import Convert from "./pages/Convert.tsx";
import Crop from "./pages/Crop.tsx";
import Compress from "./pages/Compress.tsx";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<RootLayout />}>
      <Route index element={<Home />} />
      <Route path="resize" element={<Resize />} />
      <Route path="crop" element={<Crop />} />
      <Route path="compress" element={<Compress />} />
      <Route path="convert" element={<Convert />} />
    </Route>,
  ),
);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
