import { useStoreModal } from "@/hooks/use-store-modal";
import { useEffect } from "react";

export function Root() {
  const { onOpen, isOpen, hasOpened, onClose } = useStoreModal();

  useEffect(() => {
    if (window.location.pathname !== "/") {
      onClose(true);
    }
  }, [onClose]);

  useEffect(() => {
    if (!isOpen && (!hasOpened || window.location.pathname === "/")) {
      onOpen();
    }
  }, [isOpen, hasOpened, onOpen, window.location.pathname]);

  return null;
}
