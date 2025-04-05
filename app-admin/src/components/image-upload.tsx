import { useEffect, useRef, useState } from "react";
import { Button } from "./ui/button";
import { ImagePlus, Trash } from "lucide-react";

interface ImageUploadProps {
  disabled?: boolean;
  onChange: (file: File) => void;
  onRemove: (index: number) => void;
  value: string[];
  maxImages?: number;
}

export function ImageUpload({
  disabled,
  onChange,
  onRemove,
  value,
  maxImages,
}: ImageUploadProps) {
  const [isMounted, setIsMounted] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setIsMounted(true);
  }, []);

  function handleFileChange(event: React.ChangeEvent<HTMLInputElement>) {
    const file = event.target.files?.[0];
    if (file && (!maxImages || value.length < maxImages)) {
      onChange(file);
    }
  }

  if (!isMounted) {
    return null;
  }

  return (
    <div>
      <div className="mb-4 flex items-center gap-4 flex-wrap mt-2">
        {value.map((url, index) => (
          <div
            key={url}
            className="relative w-[200px] h-[200px] rounded-md overflow-hidden border"
          >
            <Button
              type="button"
              variant="destructive"
              size="icon"
              onClick={() => onRemove(index)}
              className="absolute top-2 right-2 text-white"
            >
              <Trash className="h-4 w-4" />
            </Button>
            <img
              src={url}
              alt="Uploaded"
              className="w-full h-full object-cover"
            />
          </div>
        ))}
      </div>

      <input
        type="file"
        accept="image/*"
        ref={fileInputRef}
        className="hidden"
        onChange={handleFileChange}
      />

      <Button
        type="button"
        disabled={disabled || (maxImages ? value.length >= maxImages : false)}
        variant="secondary"
        onClick={() => fileInputRef.current?.click()}
      >
        <ImagePlus className="h-4 w-4 mr-2" /> Adicionar Imagem
      </Button>

      {maxImages && (
        <p className="text-sm text-muted-foreground mt-2">
          {value.length}/{maxImages} imagens adicionadas
        </p>
      )}
    </div>
  );
}
