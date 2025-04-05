import { getUserFirstStore } from "@/api/get-user-first-store";
import { useStoreModal } from "@/hooks/use-store-modal";
import { api } from "@/lib/axios";
import { ModalProvider } from "@/providers/modal";
import { useQuery } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { useEffect, useLayoutEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

export function RootLayout() {
  const storeModal = useStoreModal();
  const navigate = useNavigate();

  useLayoutEffect(() => {
    const interceptorId = api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (isAxiosError(error)) {
          const status = error.response?.status;
          if (status === 401) {
            storeModal.onClose();
            navigate("/login", {
              replace: false,
            });
          }
        }

        return Promise.reject(error);
      }
    );

    return () => {
      api.interceptors.response.eject(interceptorId);
    };
  }, [navigate]);

  const { data: store } = useQuery({
    queryKey: ["userFirstStore"],
    queryFn: getUserFirstStore,
    retry: false,
  });

  useEffect(() => {
    if (store) {
      if (window.location.pathname !== `/${store.id}`) {
        navigate(`/${store.id}`, { replace: true });
      }
      storeModal.onClose();
    } else {
      storeModal.onOpen();
    }
  }, [store, navigate]);

  return (
    <>
      <ModalProvider />
      <Outlet />
    </>
  );
}
