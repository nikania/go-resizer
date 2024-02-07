import { createContext, useReducer } from "react";

type FileState = {
  name: string;
};

export const FileContext = createContext({
  name: "",
});

const fileReducer = (
  state: FileState,
  action: { type: string; payload: any },
) => {
  switch (action.type) {
    case "SET_NAME":
      return { ...state, name: action.payload };
    default:
      return state;
  }
};

export function FileProvider({ children }) {
  const [state, dispatch] = useReducer(fileReducer, { name: "a1.jpg" });

  const changeName = (name: string) => {
    console.log(name);
    dispatch({ type: "SET_NAME", payload: name });
  };
  //can put custom logic here
  return (
    <FileContext.Provider value={{ ...state, changeName }}>
      {children}
    </FileContext.Provider>
  );
}
