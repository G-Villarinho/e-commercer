import { getStoreById } from "@/api/get-store-by-id";
import { Navbar } from "@/components/navbar";
import { useStoreModal } from "@/hooks/use-store-modal";
import { api } from "@/lib/axios";
import { ModalProvider } from "@/providers/modal";
import { useQuery } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { useLayoutEffect } from "react";
import { Navigate, Outlet, useNavigate, useParams } from "react-router-dom";

export function DashboardLayout() {
  const onClose = useStoreModal((state) => state.onClose);
  const { storeId } = useParams();
  const navigate = useNavigate();

  useLayoutEffect(() => {
    const interceptorId = api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (isAxiosError(error)) {
          const status = error.response?.status;
          if (status === 401) {
            onClose();
            navigate("/login", { replace: false });
          }
        }
        return Promise.reject(error);
      }
    );

    return () => {
      api.interceptors.response.eject(interceptorId);
    };
  }, [navigate, onClose]);

  if (!storeId || typeof storeId !== "string") {
    return <Navigate to="/" />;
  }

  async function getStoreByIdFn() {
    try {
      if (!storeId) {
        navigate("/", { replace: true });
      }

      return await getStoreById({ storeId: storeId as string });
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response?.status === 404) {
          navigate("/", { replace: true });
        }
      }
      throw error;
    }
  }

  const {
    data: store,
    error,
    isLoading,
  } = useQuery({
    queryKey: ["store", storeId],
    queryFn: getStoreByIdFn,
    retry: false,
    refetchOnWindowFocus: false,
    enabled: !!storeId,
  });

  if (isLoading) {
    return null;
  }

  if (error || !store) {
    return <Navigate to="/" />;
  }

  return (
    <main>
      <Navbar />
      <ModalProvider />
      <div className="flex-col">
        <div className="flex-1 space-y-4 p-8 pt-6">
          <Outlet />
        </div>
      </div>
    </main>
  );
}
