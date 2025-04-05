import { useEffect, useState } from "react";
import { Modal } from "@/components/ui/modal";
import { Button } from "@/components/ui/button";

interface AlertModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: () => void;
  loading: boolean;
}

export function AlertModal({
  isOpen,
  onClose,
  onConfirm,
  loading,
}: AlertModalProps) {
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    setIsMounted(true);
  }, []);

  if (!isMounted) {
    return null;
  }

  return (
    <Modal
      title="Você tem certeza?"
      isOpen={isOpen}
      onClose={onClose}
      description="Essa ação não pode ser desfeita."
    >
      <div className="pt-6 space-x-2 flex items-center justify-end w-full">
        <Button
          type="button"
          disabled={loading}
          variant="outline"
          onClick={onClose}
        >
          Cancelar
        </Button>
        <Button
          type="button"
          disabled={loading}
          variant="destructive"
          className="text-white"
          onClick={onConfirm}
        >
          Continuar
        </Button>
      </div>
    </Modal>
  );
}
