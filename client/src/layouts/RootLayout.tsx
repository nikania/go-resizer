import { Box, Button, Flex, Heading, Spacer } from "@chakra-ui/react";
import { Outlet, NavLink } from "react-router-dom";
import "./RootLayout.css";

export default function RootLayout() {
  return (
    <>
      <header>
        <Flex as="nav" alignItems="center" gap="10px">
          <Heading as="h1">Resizing app</Heading>
          <Box>
            <NavLink to="/">Home</NavLink>
          </Box>
          <Box>
            <NavLink to="resize">Resize</NavLink>
          </Box>
          <Box>
            <NavLink to="convert">Convert</NavLink>
          </Box>

          {/* <NavLink to="crop">Crop</NavLink>
          <NavLink to="compress">Compress</NavLink> */}
          <Spacer />
          <Button>Login</Button>
          <Button colorScheme="teal">Sign up</Button>
        </Flex>
      </header>
      <main>
        <Outlet />
      </main>
      <footer>©</footer>
    </>
  );
}
