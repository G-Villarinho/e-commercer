import { create } from "zustand";

interface useStoreModalInterface {
  isOpen: boolean;
  hasOpened: boolean;
  onOpen: () => void;
  onClose: (reset?: boolean) => void;
}

export const useStoreModal = create<useStoreModalInterface>((set) => ({
  isOpen: false,
  hasOpened: false,
  onOpen: () => set({ isOpen: true, hasOpened: true }),
  onClose: (reset = false) => set({ isOpen: false, hasOpened: reset ? false : true }),
}));
