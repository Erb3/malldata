import { createLazyFileRoute } from '@tanstack/react-router';

export const Route = createLazyFileRoute('/dashboard')({
  component: DashboardView,
});

function DashboardView() {
  return (
    <div className="flex justify-center">
      <h1>My Plots</h1>
    </div>
  );
}
