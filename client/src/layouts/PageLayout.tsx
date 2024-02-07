// page layout with common elements for every page

import React, { PropsWithChildren } from "react";
import DownloadFile from "../components/Download";
import {
  Box,
  Button,
  GridItem,
  HStack,
  SimpleGrid,
  Text,
} from "@chakra-ui/react";

// type PageLayoutProps = { form: React.ReactNode | undefined };

// const PageLayout = (props: PropsWithChildren<PageLayoutProps>) => {
const PageLayout = (props: PropsWithChildren) => {
  return (
    <>
      <Box>{props.children}</Box>
      <SimpleGrid columns={2} gap={10}>
        <Box h="200px" border="1px solid"></Box>
        <GridItem>
          <Text>items</Text>
          <HStack>
            <DownloadFile />
            <Button>Upload</Button>
          </HStack>
        </GridItem>

        {/* <UploadFile /> */}
      </SimpleGrid>
    </>
  );
};

export default PageLayout;
