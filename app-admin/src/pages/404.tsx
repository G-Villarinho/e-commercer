import fallGuy from "@/assets/fall-guy.svg";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export function NotFound() {
  const navigate = useNavigate();

  return (
    <div className="grid h-screen place-content-center bg-white px-4">
      <div className="text-center">
        <img
          className="mx-auto h-80 w-auto"
          src={fallGuy}
          alt="Illustration of a person falling"
        />

        <h1 className="mt-6 text-2xl font-bold tracking-tight text-gray-900 sm:text-4xl">
          Uh-oh!
        </h1>

        <p className="mt-4 text-gray-500">We can't find that page.</p>

        <Button onClick={() => navigate("/")}>Voltar</Button>
      </div>
    </div>
  );
}
