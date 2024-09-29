export default function SalesPage() {
  return (
    <>
      <div className="flex flex-1 p-1">
        {/* Left Section (1/3 width) */}
        <div className="w-1/3 min-w-[200px] border-r border-muted p-4">
          <div className="flex flex-col items-start gap-2 text-left">
            <p className="text-sm text-muted-foreground">Cheese Burger</p>
            <p className="text-sm text-muted-foreground">
              <b>1 Pcs</b> x Rp20.000
            </p>
          </div>
        </div>
        {/* Right Section (2/3 width) */}
        <div className="w-2/3 p-4">
          <h2 className="text-lg font-semibold">Right Section</h2>
          <p>
            This is the right section with 2/3 width. You can add relevant
            content here.
          </p>
        </div>
      </div>
    </>
  );
}
