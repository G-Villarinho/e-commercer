import { Skeleton } from "@/components/ui/skeleton";

export function CreateProductFormSkeleton() {
  return (
    <div className="space-y-8 w-full">
      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px]" />
          <Skeleton className="h-10 w-full" />
        </div>

        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px]" />
          <Skeleton className="h-10 w-full" />
        </div>

        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px]" />
          <Skeleton className="h-10 w-full" />
        </div>

        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px]" />
          <Skeleton className="h-10 w-full" />
        </div>

        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px]" />
          <div className="flex items-center gap-2">
            <Skeleton className="h-10 w-[60px]" />
            <Skeleton className="h-10 w-full" />
          </div>
        </div>

        <div />
      </div>

      <div className="flex gap-8">
        <div className="flex items-center space-x-2">
          <Skeleton className="h-5 w-5 rounded-sm" />
          <Skeleton className="h-4 w-[80px]" />
        </div>
        <div className="flex items-center space-x-2">
          <Skeleton className="h-5 w-5 rounded-sm" />
          <Skeleton className="h-4 w-[80px]" />
        </div>
      </div>

      <Skeleton className="h-10 w-[150px]" />
    </div>
  );
}
