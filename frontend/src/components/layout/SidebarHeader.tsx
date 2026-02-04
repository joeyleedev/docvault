import { useUIStore } from '../../store/uiStore';

export function SidebarHeader() {
  const { sidebarCollapsed, toggleSidebar } = useUIStore();

  return (
    <div className="flex items-center justify-between px-4 py-3 border-b border-white/20 dark:border-white/10">
      <h1 className="text-lg font-semibold tracking-tight">DocVault</h1>
      <button
        onClick={toggleSidebar}
        className="p-1.5 rounded-md hover:bg-accent/60 text-muted-foreground hover:text-foreground transition-colors duration-150"
        title="Collapse sidebar"
      >
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
          <rect width="18" height="18" x="3" y="3" rx="2" />
          <path d="M9 3v18" />
        </svg>
      </button>
    </div>
  );
}
