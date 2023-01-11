import { createContext, useState } from "react";

export const MemoryContext = createContext();

export const MemoryContextProvider = ({ children }) => {
  const [memories, setMemories] = useState(null);
  const [memory, setMemory] = useState(null);
  const [isSorted, setSorted] = useState(false);

  const value = {
    memories,
    setMemories,
    memory,
    setMemory,
    isSorted,
    setSorted,
  };

  return (
    <MemoryContext.Provider value={value}>{children}</MemoryContext.Provider>
  );
};
