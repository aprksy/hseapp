# HSE App - Product Specifications (Development Order)

## Specification Development Strategy

Following the constitution's "Shape First" principle, specifications are ordered to:
1. Define the visual foundation (App Shell) first
2. Establish core infrastructure
3. Build features incrementally
4. Enable AI/Analytics in Phase 2

---

## SPEC-001: App Shell & Visual Foundation

**Priority**: CRITICAL - Must be first  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
The app shell defines the visual identity, layout system, and foundational UI components that all features will inhabit. This is the "first-class citizen" requirement from the constitution.

### Requirements

#### 1. Layout System
- **Sidebar Navigation**: Collapsible, multi-level menu with icons
- **Top Bar**: User profile, notifications, language switcher, theme toggle
- **Main Content Area**: Responsive grid system, max-width constraints
- **Footer**: Version info, quick links, copyright

#### 2. Theme System
- **Light Theme**: Default, high contrast for office environments
- **Dark Theme**: Reduced eye strain, modern aesthetic
- **System Theme**: Auto-detect based on OS preferences
- **Theme Persistence**: LocalStorage + user profile sync

#### 3. Internationalization (i18n)
- **Languages**: Indonesian (id_ID) + English (en_ID)
- **RTL Support**: Not required initially
- **Date/Time**: Indonesian format (DD/MM/YYYY) default
- **Number Format**: Indonesian locale (1.000,00)
- **Currency**: IDR (Rp) primary, USD secondary

#### 4. Core Components (Shadcn-Svelte Blocks)
- Buttons (variants: default, outline, ghost, destructive)
- Forms (input, select, textarea, checkbox, radio)
- Data Display (tables, cards, badges, avatars)
- Feedback (alerts, toasts, progress, skeletons)
- Navigation (tabs, breadcrumbs, pagination)
- Overlays (dialogs, sheets, dropdowns, popovers)
- Charts (pie, bar, line, radar - using Svelte wrapper for charting lib)

#### 5. Responsive Breakpoints
- Mobile: < 640px (mobile-first design)
- Tablet: 640px - 1024px
- Desktop: > 1024px

#### 6. Accessibility
- WCAG 2.1 AA compliance
- Keyboard navigation support
- Screen reader compatibility
- Focus management

### Acceptance Criteria
- [ ] All Shadcn-Svelte blocks installed and themed
- [ ] Theme toggle works with persistence
- [ ] Language switcher changes all UI text instantly
- [ ] Layout responsive across all breakpoints
- [ ] Color contrast passes WCAG AA
- [ ] Component storybook/examples documented

### Technical Notes
- Use `svelte-i18n` or similar for i18n
- CSS variables for theming (HSL color space recommended)
- TailwindCSS for utility classes (Shadcn standard)

---

## SPEC-002: Authentication & Authorization System

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
Multi-tenant authentication with role-based access control, supporting simple login, OAuth, and optional MFA.

### Requirements

#### 1. Authentication Methods
- **Email/Password**: Standard credentials with validation
- **OAuth Providers**: Google, Microsoft, GitHub (configurable per tenant)
- **MFA**: TOTP-based (Google Authenticator compatible), optional per tenant
- **Session Management**: JWT tokens, refresh token rotation

#### 2. Multi-Tenancy
- **Tenant Isolation**: Database-level row security policies (RLS)
- **Tenant Onboarding**: Self-service signup or admin provisioning
- **Custom Domains**: Optional subdomain per tenant (future)
- **Branding**: Tenant-specific logos/colors (future)

#### 3. Role-Based Access Control (RBAC)
- **Predefined Roles**:
  - Super Admin (platform level)
  - Tenant Admin (organization level)
  - HSE Manager
  - Site Supervisor
  - Safety Officer
  - Field Worker
  - Viewer/Auditor (read-only)
- **Custom Roles**: Tenant-admin can create custom roles with granular permissions
- **Permission Types**: Create, Read, Update, Delete, Approve, Export

#### 4. User Management
- **Profile**: Name, email, phone, avatar, timezone, language preference
- **Multi-Site Access**: Users can belong to multiple sites within tenant
- **Site Roles**: Different roles per site (e.g., Manager at Site A, Viewer at Site B)
- **User Status**: Active, Inactive, Suspended, Pending Activation

#### 5. Security Features
- **Password Policy**: Configurable complexity, expiration, history
- **Account Lockout**: After N failed attempts
- **Session Timeout**: Configurable idle timeout
- **Audit Log**: All auth events logged (login, logout, password change, etc.)

### Acceptance Criteria
- [ ] Email/password registration and login functional
- [ ] OAuth flow working for at least Google provider
- [ ] MFA enrollment and verification working
- [ ] RLS policies enforce tenant isolation
- [ ] Role-based UI element visibility (buttons, menus)
- [ ] Multi-site user assignment working
- [ ] Audit logs capture all auth events

### Technical Notes
- Supabase Auth for authentication backend
- Custom claims in JWT for role/site information
- Supabase RLS for data isolation

---

## SPEC-003: Multi-Tenant & Multi-Site Architecture

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
Data model and architecture supporting multiple organizations (tenants) with multiple operational sites each.

### Requirements

#### 1. Tenant Model
- **Tenant Entity**: ID, name, slug, status, subscription tier, settings
- **Tenant Settings**:
  - Default language
  - Default timezone
  - Date/number formats
  - Logo/branding assets
  - Feature flags (MFA, OAuth providers, etc.)
  - Compliance frameworks enabled

#### 2. Site Model
- **Site Entity**: ID, tenant_id, name, code, address, GPS coordinates, status
- **Site Properties**:
  - Site type (office, factory, construction, warehouse, etc.)
  - Operating hours
  - Maximum capacity
  - Site-specific contacts
  - Site-specific regulations applicable

#### 3. Data Isolation
- **Row Level Security**: All queries filtered by tenant_id automatically
- **Cross-Tenant Access**: Explicitly prohibited unless super admin
- **Site Filtering**: Users see only assigned sites
- **Data Export**: Exports include tenant/site metadata

#### 4. Hierarchy & Relationships
```
Tenant (1) → (N) Sites
Tenant (1) → (N) Users
User (N) → (N) Sites (through user_site_roles)
Site (1) → (N) Assets/Equipment
Site (1) → (N) Incidents/Inspections/etc.
```

### Acceptance Criteria
- [ ] Tenant creation and configuration working
- [ ] Site CRUD operations functional
- [ ] RLS policies prevent cross-tenant data access
- [ ] User can be assigned to multiple sites
- [ ] Site-specific data filtering works in queries
- [ ] Tenant settings affect UI behavior (language, formats)

### Technical Notes
- Supabase RLS policies on every table
- tenant_id in JWT claims for automatic filtering
- Consider citus extension for horizontal scaling (future)

---

## SPEC-004: Approval Routing & Workflow Engine

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
Customizable approval workflows for incidents, permits, corrective actions, and other HSE processes.

### Requirements

#### 1. Workflow Definition
- **Workflow Templates**: Pre-built templates for common processes
- **Custom Workflows**: Admin can define custom approval chains
- **Step Types**:
  - Approval (single approver)
  - Parallel Approval (multiple approvers, any can approve)
  - Sequential Approval (multiple approvers, ordered)
  - Notification (info only, no action required)
  - Condition (branching based on data values)

#### 2. Routing Rules
- **Amount-Based**: Different approvers based on severity/cost thresholds
- **Role-Based**: Route to specific roles (e.g., "Site Supervisor")
- **User-Based**: Route to specific individuals
- **Department-Based**: Route to department heads
- **Site-Based**: Route within same site hierarchy

#### 3. Approval Actions
- **Approve**: Move to next step or complete workflow
- **Reject**: Return to submitter with comments
- **Request Changes**: Return for modification, re-submit to same step
- **Delegate**: Assign to another approver temporarily
- **Comment**: Add notes without changing status

#### 4. Notifications
- **Email Notifications**: On assignment, approval, rejection, reminder
- **In-App Notifications**: Real-time via Supabase real-time
- **Push Notifications**: Mobile push via FCM
- **Reminder Escalation**: Auto-escalate if not acted upon in N hours

#### 5. Audit Trail
- Complete history of all actions
- Timestamp, actor, action, comments
- Before/after snapshots for data changes

### Acceptance Criteria
- [ ] Workflow template creation UI functional
- [ ] Custom workflow definition working
- [ ] Approval routing follows defined rules
- [ ] Notifications sent at appropriate times
- [ ] Audit trail captures all workflow actions
- [ ] Delegation feature working
- [ ] Escalation rules functional

### Technical Notes
- State machine pattern for workflow states
- Supabase database triggers for state transitions
- Redis queue for notification processing (optional)

---

## SPEC-005: Compliance Management Module

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
Core HSE functionality for tracking regulatory compliance, scoring, indicators, and checklists aligned with Indonesian standards.

### Requirements

#### 1. Regulatory Framework
- **Indonesian Regulations Library**:
  - UU No. 1 Tahun 1970 (Keselamatan Kerja)
  - Permenaker related to K3
  - KLHK environmental regulations
  - Industry-specific regulations
- **International Standards** (optional):
  - ISO 45001 (Occupational Health & Safety)
  - ISO 14001 (Environmental Management)
  - SMKP (Mining sector)
  - PROPER (Environmental rating)

#### 2. Compliance Scoring
- **Score Calculation**: Weighted average of checklist items
- **Score Visualization**: Gauge charts, trend lines
- **Threshold Alerts**: Warning when score drops below threshold
- **Benchmarking**: Compare against industry averages (future)

#### 3. Compliance Indicators
- **Leading Indicators**: Training completion, inspection frequency, near-miss reports
- **Lagging Indicators**: Incident rates, violations, downtime
- **KPI Dashboard**: Visual display of key metrics
- **Target Setting**: Define targets per indicator

#### 4. Checklists
- **Checklist Templates**: Pre-built per regulation/standard
- **Custom Checklists**: Tenant can create own checklists
- **Checklist Items**: Yes/No/NA with evidence requirements
- **Scoring**: Points per item, weighted sections
- **Frequency**: One-time, daily, weekly, monthly, quarterly, annually
- **Assignment**: Assign to roles/users/sites

#### 5. Evidence Collection
- **Photo Upload**: Camera capture or file upload
- **Document Attachment**: PDFs, spreadsheets
- **Signatures**: Digital signature capture
- **Geotagging**: GPS coordinates for field inspections
- **Timestamp**: Automatic recording

### Acceptance Criteria
- [ ] Regulation library browseable and searchable
- [ ] Compliance scores calculated correctly
- [ ] Indicator dashboard displays real-time data
- [ ] Checklists can be created and assigned
- [ ] Checklist execution captures evidence
- [ ] Scores update automatically upon checklist completion
- [ ] Alerts trigger when thresholds breached

### Technical Notes
- Supabase storage for evidence files
- Calculated columns for compliance scores
- Materialized views for performance on dashboards

---

## SPEC-006: Forms & Data Capture (Web)

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
Flexible form builder and data capture interface for HSE documentation on web platform.

### Requirements

#### 1. Form Builder
- **Drag-and-Drop Interface**: Visual form construction
- **Field Types**:
  - Text (single line, multi-line)
  - Number (integer, decimal)
  - Date/Time (date picker, time picker, datetime)
  - Select (dropdown, multi-select, searchable)
  - Boolean (checkbox, toggle)
  - File (image, document, multiple files)
  - Signature (canvas for digital signature)
  - Location (GPS coordinates, map picker)
  - Table/Grid (repeating rows)
- **Validation Rules**: Required, min/max, regex patterns, custom logic
- **Conditional Logic**: Show/hide fields based on previous answers
- **Sections & Pages**: Organize complex forms

#### 2. Form Templates
- **Pre-Built Templates**: Incident report, inspection, permit, audit, etc.
- **Template Library**: Browse and duplicate templates
- **Version Control**: Track form template versions
- **Publishing**: Draft → Published → Archived lifecycle

#### 3. Form Submission
- **Auto-Save**: Draft preservation during completion
- **Offline Support**: Queue submissions when offline (with service workers)
- **Validation**: Client-side and server-side validation
- **Submission Tracking**: Status tracking (draft, submitted, approved, rejected)
- **Duplicate Detection**: Warn on potential duplicate submissions

#### 4. Data Entry Features
- **Bulk Edit**: Update multiple records simultaneously
- **Import from Spreadsheet**: CSV/Excel import with mapping
- **Copy Previous**: Duplicate prior submission for editing
- **Quick Actions**: Common responses as shortcuts

### Acceptance Criteria
- [ ] Form builder UI intuitive and functional
- [ ] All field types render correctly
- [ ] Validation rules enforced client and server-side
- [ ] Conditional logic shows/hides fields appropriately
- [ ] Form templates can be created and published
- [ ] Submissions save correctly with all data
- [ ] Import from Excel/CSV working with field mapping
- [ ] Offline form completion queues for sync

### Technical Notes
- JSON schema for form definition
- Reactive form validation with Svelte stores
- IndexedDB for offline form drafts

---

## SPEC-007: Mobile Data Capture (Flutter Spec)

**Priority**: HIGH  
**Phase**: 1 (Parallel with Web Forms)  
**Status**: Pending Approval  

### Overview
Mobile application specification for field data capture, optimized for offline use and camera/GPS capabilities.

*Note: This spec defines requirements for the separate Flutter mobile project.*

### Requirements

#### 1. Offline-First Architecture
- **Local Database**: Hive or SQLite for local data storage
- **Sync Engine**: Background sync with conflict resolution
- **Queue Management**: Outbound submission queue
- **Conflict Resolution**: Last-write-wins or manual resolution
- **Sync Status**: Visual indicator of sync state

#### 2. Camera Integration
- **Photo Capture**: In-app camera with preview
- **Multiple Photos**: Batch photo capture
- **Video Recording**: Short video clips (optional)
- **Image Annotation**: Draw on photos, add arrows/text
- **Compression**: Optimize image size before upload
- **EXIF Data**: Preserve timestamp, GPS coordinates

#### 3. GPS & Location Services
- **Current Location**: Auto-capture GPS coordinates
- **Accuracy Indicator**: Show GPS accuracy level
- **Map View**: Display location on map
- **Geofencing**: Validate location against site boundaries
- **Offline Maps**: Cached map tiles for remote areas

#### 4. Mobile-Optimized Forms
- **Responsive Layout**: Adapt to phone/tablet screens
- **Touch-Friendly**: Large tap targets, swipe gestures
- **Voice Input**: Speech-to-text for text fields (optional)
- **Barcode/QR Scanner**: Scan equipment tags, permits
- **Signature Pad**: Touch-based signature capture

#### 5. Push Notifications
- **FCM Integration**: Firebase Cloud Messaging
- **Notification Types**: Assignment, approval request, reminder, alert
- **Deep Linking**: Tap notification opens relevant screen
- **Badge Count**: Unread notification count on app icon

#### 6. Mobile UX Patterns
- **Bottom Navigation**: Primary navigation pattern
- **Pull-to-Refresh**: Refresh data lists
- **Swipe Actions**: Quick actions on list items
- **Haptic Feedback**: Tactile response for actions
- **Dark Mode**: Match system theme

### Acceptance Criteria
- [ ] App functions fully offline for core tasks
- [ ] Sync completes successfully when online
- [ ] Camera capture integrates with forms
- [ ] GPS coordinates captured accurately
- [ ] Forms render correctly on mobile screens
- [ ] Push notifications received and actionable
- [ ] App handles background/foreground transitions

### Technical Notes
- Flutter 3.x with null safety
- Riverpod or Bloc for state management
- Dio for HTTP client with retry logic
- Workmanager for background sync tasks

---

## SPEC-008: Reporting & Export Engine

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Requirements

#### 1. Report Types
- **Compliance Reports**: Score summaries, checklist completion
- **Incident Reports**: Detailed incident analysis
- **Inspection Reports**: Findings and recommendations
- **Audit Reports**: Audit trails and findings
- **Executive Dashboards**: High-level KPI summaries
- **Regulatory Reports**: Format per Indonesian regulatory requirements

#### 2. Export Formats
- **PDF**: Professional formatting, headers/footers, page numbers
- **PNG**: Image export of charts/dashboards
- **Spreadsheet**: Excel (.xlsx) and CSV
- **TXT**: Plain text export
- **Word**: DOCX format (optional, future)

#### 3. Report Builder
- **Drag-and-Drop**: Visual report layout designer
- **Data Sources**: Select from available data entities
- **Filters**: Date ranges, sites, departments, etc.
- **Grouping**: Group data by categories
- **Calculations**: Sum, average, count, custom formulas
- **Charts**: Embed pie, bar, line, radar charts

#### 4. Scheduling & Distribution
- **Scheduled Reports**: Daily, weekly, monthly automation
- **Email Delivery**: Send reports to stakeholders
- **Report Library**: Save and share report templates
- **Access Control**: Restrict report access by role

#### 5. Customization
- **Branding**: Tenant logo, colors, custom headers
- **Templates**: Pre-designed templates per report type
- **Localization**: Reports in selected language
- **Watermarks**: Draft/confidential watermarks

### Acceptance Criteria
- [ ] All export formats generate correctly
- [ ] PDF formatting professional and readable
- [ ] Excel exports preserve data types and formatting
- [ ] Report builder allows custom report creation
- [ ] Scheduled reports generate and send automatically
- [ ] Charts render correctly in exports
- [ ] Localization applies to report content

### Technical Notes
- Puppeteer or Playwright for PDF generation (HTML to PDF)
- SheetJS or Go excel library for spreadsheet export
- Cron jobs or Supabase Edge Functions for scheduling

---

## SPEC-009: Dashboards & Visualizations

**Priority**: HIGH  
**Phase**: 1  
**Status**: Pending Approval  

### Requirements

#### 1. Chart Types
- **Pie Charts**: Proportions, percentages
- **Bar Charts**: Comparisons across categories
- **Line Charts**: Trends over time
- **Radar/Spider Charts**: Multi-dimensional comparisons
- **Gauge Charts**: Single metric against target
- **Heat Maps**: Density/intensity visualization

#### 2. Dashboard Features
- **Widget-Based**: Drag-and-drop widget arrangement
- **Resizeable**: Adjust widget sizes
- **Filtering**: Global filters apply to all widgets
- **Drill-Down**: Click to see detailed data
- **Real-Time Updates**: Live data via Supabase real-time
- **Export**: Export dashboard as image or PDF

#### 3. Pre-Built Dashboards
- **Executive Dashboard**: High-level KPIs, compliance score
- **HSE Manager Dashboard**: Open actions, trends, alerts
- **Site Dashboard**: Site-specific metrics
- **Compliance Dashboard**: Regulatory compliance status
- **Incident Dashboard**: Incident statistics and trends

#### 4. Customization
- **Create Custom Dashboard**: User-defined dashboards
- **Save Layout**: Persist widget arrangement
- **Share Dashboard**: Share with team members
- **Theme-Aware**: Charts adapt to light/dark theme

#### 5. Performance
- **Lazy Loading**: Load visible widgets first
- **Caching**: Cache query results
- **Pagination**: Paginate large datasets
- **Aggregation**: Pre-aggregate data where possible

### Acceptance Criteria
- [ ] All chart types render correctly
- [ ] Dashboard widgets can be rearranged
- [ ] Filters apply across all widgets
- [ ] Real-time updates reflect immediately
- [ ] Pre-built dashboards provide value out-of-box
- [ ] Custom dashboard creation intuitive
- [ ] Charts respect theme colors
- [ ] Performance acceptable with large datasets

### Technical Notes
- Chart.js or ApexCharts for visualizations
- Svelte stores for reactive dashboard state
- Virtual scrolling for large datasets

---

## SPEC-010: Project Management (Simple)

**Priority**: MEDIUM  
**Phase**: 1  
**Status**: Pending Approval  

### Requirements

#### 1. Task Management
- **Task Creation**: Title, description, assignee, due date, priority
- **Task Status**: To Do, In Progress, Review, Done (customizable)
- **Subtasks**: Break down tasks into smaller items
- **Checklists**: Item-level completion tracking
- **Attachments**: Files, images linked to tasks
- **Comments**: Discussion thread per task

#### 2. Kanban Board
- **Board View**: Columns represent status/stages
- **Card Design**: Task summary on draggable cards
- **Drag-and-Drop**: Move tasks between columns
- **WIP Limits**: Optional work-in-progress limits per column
- **Swimlanes**: Horizontal grouping (by assignee, priority, etc.)

#### 3. Timeline View
- **Gantt Chart**: Visual timeline of tasks
- **Dependencies**: Link tasks (finish-to-start, etc.)
- **Milestones**: Key dates/markers
- **Critical Path**: Highlight critical tasks (optional)
- **Baseline**: Compare actual vs. planned (optional)

#### 4. Assignment & Monitoring
- **Assignee Selection**: Assign to individuals or teams
- **Workload View**: See assignee's task load
- **Overdue Tracking**: Highlight overdue tasks
- **Time Tracking**: Log time spent on tasks (optional)
- **Reminders**: Due date reminders via notification

#### 5. Project Templates
- **Pre-Built Templates**: Common HSE project types
- **Custom Templates**: Save own templates
- **Clone Projects**: Duplicate existing projects
- **Task Templates**: Reusable task structures

### Acceptance Criteria
- [ ] Task CRUD operations functional
- [ ] Kanban board drag-and-drop smooth
- [ ] Timeline view displays correctly
- [ ] Assignments notify users
- [ ] Comments and attachments work
- [ ] Templates can be created and applied
- [ ] Overdue tasks highlighted appropriately

### Technical Notes
- dnd-kit or similar for drag-and-drop
- Supabase real-time for collaborative boards
- Consider vis-timeline or similar for Gantt

---

## SPEC-011: Import/Export (Common Formats)

**Priority**: MEDIUM  
**Phase**: 1  
**Status**: Pending Approval  

### Requirements

#### 1. Import Formats
- **CSV**: Comma-separated values
- **Excel**: .xlsx, .xls formats
- **JSON**: Structured data import

#### 2. Import Features
- **Field Mapping**: Map import columns to database fields
- **Data Validation**: Validate data before import
- **Error Handling**: Report errors with row numbers
- **Preview**: Preview data before committing
- **Batch Processing**: Handle large files in batches
- **Rollback**: Undo import if errors detected

#### 3. Export Formats
- **CSV**: Universal format
- **Excel**: .xlsx with formatting
- **JSON**: API-style exports
- **PDF**: Formatted documents
- **PNG**: Chart/dashboard images

#### 4. Bulk Operations
- **Bulk Update**: Update records via import
- **Bulk Delete**: Delete multiple records
- **Bulk Export**: Export filtered datasets
- **Background Processing**: Long operations run async

### Acceptance Criteria
- [ ] CSV import with field mapping working
- [ ] Excel import preserves data types
- [ ] Import validation catches errors
- [ ] Error reports clear and actionable
- [ ] All export formats generate correctly
- [ ] Large file handling doesn't timeout
- [ ] Background jobs show progress

---

## SPEC-012: Notifications & Communication

**Priority**: MEDIUM  
**Phase**: 1  
**Status**: Pending Approval  

### Overview
Comprehensive notification system supporting two distinct categories: **Activity Notifications** (routine workflow) and **Alerts** (critical incidents). Includes support for cross-site teams and module-specific routing.

### Requirements

#### 1. Module Architecture & Planning Loop
The notification system shall support the HSE management cycle with clear separation:

**Modules:**
- **Planning**: Risk assessments, permit applications, inspection schedules, training plans
- **Monitoring**: Real-time observations, checklist executions, sensor data, patrol logs
- **Evaluation**: Incident investigations, audit findings, compliance scoring, performance reviews
- **Action/Improvement**: Corrective actions, preventive actions, lessons learned

**The Loop:**
```
Planning → Monitoring → Evaluation → Action → (back to) Planning
```

**Override Mechanisms:**
- **Escalation Override**: Bypass normal routing when SLA breached
- **Emergency Override**: Direct notify responsible parties regardless of hierarchy
- **Delegation Override**: Temporary reassignment during absences
- **Priority Override**: High-severity items skip queue

#### 2. Cross-Site Team Support
Teams/groups can span across multiple sites within a tenant:

- **Cross-Site Roles**: 
  - Regional/Group Managers oversee multiple sites
  - Subject Matter Experts (SMEs) support all sites
  - Shared Services (HR, Legal, Finance) serve entire tenant
  
- **Team Structures**:
  - **Site-Based Team**: All members at single site
  - **Multi-Site Team**: Members distributed across sites (e.g., all Site Supervisors)
  - **Functional Team**: Role-based grouping (e.g., all Safety Officers)
  - **Project Team**: Temporary team for specific initiative spanning sites
  
- **Notification Routing**:
  - Route to team → all members receive notification
  - Route to role → users with that role at relevant site(s)
  - Route to individual → specific user regardless of site
  - Route to site → current site leadership team

#### 3. Notification Categories

**Category A: Activity Notifications (Routine Workflow)**
*Lower urgency, workflow-driven, actionable items*

| Type | Trigger Example | Recipients | Actions Available |
|------|----------------|------------|-------------------|
| **Task Assignment** | Lead assigns inspection task to team member | Assignee, Task Observer | Accept, Decline, Request Clarification, Delegate |
| **Task Received** | Staff receives newly assigned task | Assignee | View Details, Add to Calendar, Acknowledge |
| **Report Submission** | Site submits monthly HSE report | Site Manager, Regional HSE Manager, Compliance Officer | Review, Approve, Request Revision, Forward |
| **Approval Query** | Permit-to-work submitted for approval | Assigned Approver, Backup Approver | Approve, Reject, Request Changes, Delegate, Comment |
| **Comment/Mention** | User @mentioned in incident discussion | Mentioned User, Thread Participants | Reply, View Context, Dismiss |
| **Status Change** | Corrective action status updated | Originator, Assignee, Stakeholders | View Update, Add Comment, Escalate |
| **Reminder** | Checklist due in 2 days | Responsible Person, Supervisor | Complete Now, Snooze, Reassign |
| **Completion Notice** | Training completed by team member | Team Lead, HSE Manager, HR | View Certificate, Record in Profile, Acknowledge |

**Category B: Alerts (Critical/Urgent)**
*High urgency, immediate attention required, potential safety impact*

| Type | Trigger Example | Recipients | Actions Available |
|------|----------------|------------|-------------------|
| **Accident/Incident** | Fatal/near-miss/high-potential incident reported | Site Manager, HSE Manager, Emergency Response Team, Regional Director, Corporate HSE | Acknowledge, Mobilize Team, Initiate Investigation, Notify Authorities, Escalate Further |
| **Environmental Breach** | Emission/exceedance detected | Environmental Officer, Site Manager, KLHK Liaison | Investigate, Mitigate, Report to Authority, Document |
| **Regulatory Deadline** | Compliance submission due tomorrow | Compliance Officer, HSE Manager, Legal | Submit Now, Request Extension, Escalate |
| **Threshold Breach** | Safety score drops below minimum | HSE Manager, Site Supervisor, Regional Manager | Review Root Cause, Implement Controls, Schedule Audit |
| **Equipment Failure** | Critical safety equipment malfunction | Maintenance Lead, Safety Officer, Operations Manager | Isolate Equipment, Arrange Repair, Issue Work Stoppage |
| **Evacuation/Emergency** | Emergency drill or real evacuation | All Personnel On-Site | Acknowledge, Proceed to Assembly Point, Account for Team |
| **Medical Emergency** | Injury requiring medical attention | First Aid Team, Site Medic, Emergency Contacts | Respond, Call Ambulance, Document Incident |
| **Stop Work Order** | Authority issues stop work directive | Site Manager, All Affected Teams, Contractor Leads | Acknowledge, Halt Operations, Communicate to Crew |

#### 4. Notification Channels
- **In-App**: Bell icon with notification list, real-time via Supabase Realtime
- **Email**: HTML emails with tenant branding, deep links to relevant records
- **Push**: Mobile push via FCM with category-specific sounds/badges
- **SMS**: Optional SMS for Alerts only (Twilio integration)
- **Webhook**: Outbound webhook for third-party integrations (Slack, Teams)

#### 5. Notification Metadata Structure
```json
{
  "id": "uuid",
  "tenant_id": "uuid",
  "category": "activity|alert",
  "type": "task_assignment|incident_report|etc",
  "module": "planning|monitoring|evaluation|action",
  "priority": "low|normal|high|urgent|critical",
  "title": "string",
  "body": "string",
  "actor": { "user_id": "uuid", "name": "string", "role": "string" },
  "recipients": [{ "user_id": "uuid", "site_id": "uuid?" }],
  "context": {
    "entity_type": "incident|task|permit|etc",
    "entity_id": "uuid",
    "site_id": "uuid?",
    "previous_status": "string?",
    "new_status": "string?"
  },
  "actions": ["acknowledge", "approve", "view", "etc"],
  "channels": ["in_app", "email", "push", "sms"],
  "created_at": "timestamp",
  "read_at": "timestamp?",
  "acted_at": "timestamp?"
}
```

#### 6. User Preferences
- **Channel Selection**: Choose channels per notification type/category
- **Frequency**: 
  - Immediate (real-time)
  - Batched (every 15min, hourly)
  - Daily digest (summary email)
  - Weekly digest (comprehensive report)
- **Quiet Hours**: Suppress non-alert notifications during specified times (e.g., 22:00-06:00)
- **Site Filtering**: Only receive notifications for assigned sites
- **Role-Based Defaults**: Pre-configured based on user role
- **Unsubscribe**: Opt-out of non-critical activity notifications (alerts cannot be disabled)

#### 7. Notification Center (UI)
- **Unified Inbox**: Tabbed view (All, Activity, Alerts, Unread)
- **Quick Filters**: By module (Planning, Monitoring, Evaluation, Action)
- **Search**: Full-text search across notification content
- **Bulk Actions**: Mark all as read, archive, delete
- **Deep Linking**: Click notification → navigate to relevant record/form
- **Action Buttons**: Contextual actions inline (Approve/Reject without leaving inbox)
- **Snooze**: Temporarily hide notification, reappear later
- **Escalation Path**: Visual indicator of escalation chain

#### 8. Delivery & Retry Logic
- **Immediate Delivery**: Alerts sent instantly via all configured channels
- **Batching**: Activity notifications batched per user preference
- **Retry Policy**: 3 retries with exponential backoff for failed deliveries
- **Fallback**: If push fails, escalate to email; if email fails, try SMS (for alerts)
- **Delivery Receipt**: Track delivery status per channel

#### 9. Audit & Analytics
- **Delivery Logs**: Timestamp, channel, status (sent/delivered/failed/read/acted)
- **Response Time**: Time from delivery to user action
- **Engagement Metrics**: Open rates, action rates by notification type
- **SLA Tracking**: Monitor response times against SLA targets
- **Escalation History**: Track when and why escalations occurred

### Acceptance Criteria
- [ ] Activity notifications and alerts clearly distinguished in UI (color, icon, sound)
- [ ] Module separation (Planning/Monitoring/Evaluation/Action) implemented in data model
- [ ] Cross-site team notification routing working correctly
- [ ] All activity notification types listed above functional with correct actions
- [ ] All alert types listed above functional with correct actions
- [ ] Override mechanisms (escalation, emergency, delegation) working
- [ ] User preferences respected across all channels
- [ ] Notification center supports filtering by module and category
- [ ] Deep linking navigates to correct record with context
- [ ] Inline actions (approve/reject) execute without page navigation
- [ ] Quiet hours suppress non-alert notifications
- [ ] Alert notifications bypass quiet hours and reach user via fallback channels
- [ ] Delivery retry logic functional with proper logging
- [ ] Audit trail captures all notification events and user responses
- [ ] Email templates professional with tenant branding
- [ ] Push notifications display correctly on iOS and Android with category badges

### Technical Notes
- Use Supabase Realtime for instant in-app notifications
- Firebase Cloud Messaging (FCM) for mobile push
- Go background workers for email/SMS processing (use `gomail` or similar)
- Redis queue for notification job processing with priority queues
- Store notification templates in database with i18n support (id_ID, en_ID)
- Implement idempotency keys to prevent duplicate notifications
- Use PostgreSQL NOTIFY/LISTEN for real-time events
- Consider rate limiting to prevent notification spam

---

## SPEC-013: AI Chat & Analysis (BYOK)

**Priority**: LOW  
**Phase**: 2 (Post-MVP)  
**Status**: Future Consideration  

### Overview
AI-powered chatbot, analysis, and recommendation engine using Bring Your Own Key (BYOK) model for LLM APIs.

*This spec is placeholder for Phase 2 planning.*

### Requirements (High-Level)
- **Chat Interface**: Conversational UI for Q&A
- **LLM Integration**: OpenAI, Anthropic, or local models
- **Context Awareness**: Understand HSE domain
- **Recommendations**: Suggest actions based on data patterns
- **Analysis**: Natural language queries on data
- **BYOK**: Users provide their own API keys

### Deferral Rationale
- Requires mature data foundation
- Model training/fine-tuning needed
- Cost considerations for LLM usage
- Phase 1 focus on core HSE functionality

---

## SPEC-014: Settings & Configuration

**Priority**: MEDIUM  
**Phase**: 1  
**Status**: Pending Approval  

### Requirements

#### 1. Tenant Settings
- Profile (name, logo, contact info)
- Localization defaults
- Feature flags
- Subscription/billing (if applicable)

#### 2. User Settings
- Profile management
- Password change
- MFA enrollment
- Notification preferences
- Theme/language preferences

#### 3. Site Settings
- Site details
- Operating hours
- Site-specific regulations
- Site personnel

#### 4. System Configuration
- Email server settings
- OAuth provider configuration
- Backup/export settings
- Audit log retention

### Acceptance Criteria
- [ ] All settings categories accessible
- [ ] Changes persist correctly
- [ ] Validation prevents invalid configurations
- [ ] Settings respect role-based access

---

## Development Order Summary

| Order | Spec ID | Name | Phase | Priority |
|-------|---------|------|-------|----------|
| 1 | SPEC-001 | App Shell & Visual Foundation | 1 | CRITICAL |
| 2 | SPEC-002 | Authentication & Authorization | 1 | HIGH |
| 3 | SPEC-003 | Multi-Tenant & Multi-Site | 1 | HIGH |
| 4 | SPEC-004 | Approval Routing | 1 | HIGH |
| 5 | SPEC-005 | Compliance Management | 1 | HIGH |
| 6 | SPEC-006 | Forms & Data Capture (Web) | 1 | HIGH |
| 7 | SPEC-007 | Mobile Data Capture (Flutter) | 1 | HIGH |
| 8 | SPEC-008 | Reporting & Export | 1 | HIGH |
| 9 | SPEC-009 | Dashboards & Visualizations | 1 | HIGH |
| 10 | SPEC-010 | Project Management (Simple) | 1 | MEDIUM |
| 11 | SPEC-011 | Import/Export | 1 | MEDIUM |
| 12 | SPEC-012 | Notifications | 1 | MEDIUM |
| 13 | SPEC-014 | Settings & Configuration | 1 | MEDIUM |
| 14 | SPEC-013 | AI Chat & Analysis | 2 | LOW |

---

**Document Status**: Draft v1.0  
**Created**: Initial Specifications  
**Next Step**: Approval required before implementation planning  
**Approval Gate**: Each spec requires approval before corresponding implementation
