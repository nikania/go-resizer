// page layout with common elements for every page

import { PropsWithChildren } from "react";
import { Box, GridItem, SimpleGrid, Text } from "@chakra-ui/react";
import UploadFile from "../components/Upload";

// type PageLayoutProps = { form: React.ReactNode | undefined };

// const PageLayout = (props: PropsWithChildren<PageLayoutProps>) => {
const PageLayout = (props: PropsWithChildren) => {
  return (
    <>
      <Box>{props.children}</Box>
      <SimpleGrid columns={2} gap={10}>
        <Box h="200px" border="1px solid">
          <UploadFile />
        </Box>
        <GridItem>
          <Text>items</Text>
        </GridItem>
      </SimpleGrid>
    </>
  );
};

export default PageLayout;
