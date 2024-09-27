import { Button } from "@/components/ui/button";

export const description =
  "A products dashboard with a sidebar navigation and a main content area. The dashboard has a header with a search input and a user menu. The sidebar has a logo, navigation links, and a card with a call to action. The main content area shows an empty state with a call to action.";

export default function SalesPage() {
  return (
    <>
      <div className="flex flex-1 p-4">
        {/* Left Section (1/3 width) */}
        <div className="flex-2 border-r border-muted p-4">
          <div className="flex flex-col items-center gap-1 text-center">
            <h3 className="text-2xl font-bold tracking-tight">
              You have no sales
            </h3>
            <p className="text-sm text-muted-foreground">
              You can start selling as soon as you add a sale.
            </p>
            <Button className="mt-4">Add Sales</Button>
          </div>
        </div>
        {/* Right Section (2/3 width) */}
        <div className="flex-1 p-4">
          <h2 className="text-lg font-semibold">Left Section</h2>
          <p>
            This is the left section with 1/3 width. You can add relevant
            content here.
          </p>
        </div>
      </div>
    </>
  );
}
