package badgetemplates

// summaryTemplate is the SVG template for the Summary badge style (400x200)
const DefaultSummary = `<svg
  width="400"
  height="200"
  viewBox="0 0 400 200"
  xmlns="http://www.w3.org/2000/svg"
  role="img"
  aria-label="Open Source Contribution Summary">
  <defs>
    <style>
      text {
        font-family: system-ui, -apple-system, BlinkMacSystemFont,
                     "Segoe UI", Helvetica, Arial, sans-serif;
      }
      .bg {
        fill: {{.Colors.Background}};
      }
      .card {
        fill: {{.Colors.BackgroundAlt}};
      }
      .username {
        font-size: 18px;
        font-style: italic;
        font-weight: 700;
        fill: {{.Colors.Text}};
      }
      .subtitle {
        font-size: 12px;
        fill: {{.Colors.TextSecondary}};
      }
      .stat-value {
        font-size: 26px;
        font-style: italic;
        font-weight: 700;
        fill: {{.Colors.Text}};
      }
      .stat-label {
        font-size: 11px;
        fill: {{.Colors.TextSecondary}};
      }
    </style>
  </defs>
  <!-- Background -->
  <rect class="bg" width="400" height="200" rx="16"/>
  <!-- Header -->
  <text class="username" x="28" y="42">@{{.Stats.Username}}</text>
  <text class="subtitle" x="28" y="62">open source contributions</text>
  <!-- Stat Cards -->
  <rect class="card" x="22" y="91" width="108" height="70" rx="10"/>
  <rect class="card" x="146" y="91" width="108" height="70" rx="10"/>
  <rect class="card" x="270" y="91" width="108" height="70" rx="10"/>
  <!-- Stats -->
  <text class="stat-value" x="76" y="123" text-anchor="middle">{{.TotalProjects}}</text>
  <text class="stat-label" x="76" y="144" text-anchor="middle">PROJECTS</text>
  <text class="stat-value" x="200" y="123" text-anchor="middle">{{.TotalPRs}}</text>
  <text class="stat-label" x="200" y="144" text-anchor="middle">PRS MERGED</text>
  <text class="stat-value" x="324" y="123" text-anchor="middle">{{.TotalLines}}</text>
  <text class="stat-label" x="324" y="144" text-anchor="middle">LINES CHANGED</text>
</svg>
`

// compactTemplate is the SVG template for the Compact badge style (280x28) - Shields.io style
const DefaultCompact = `<svg
  width="280"
  height="32"
  viewBox="0 0 280 32"
  xmlns="http://www.w3.org/2000/svg"
  role="img"
  aria-label="OSS Contributions">

  <defs>
    <style>
      text {
        font-family: system-ui, -apple-system, BlinkMacSystemFont,
                     "Segoe UI", Helvetica, Arial, sans-serif;
        fill: {{.Colors.Text}};
        font-size: 12px;
        font-weight: 700;
      }

      .card {
        fill: {{.Colors.Background}};
        stroke: {{.Colors.Border}};
        stroke-width: 1;
        rx: 16;
      }
    </style>

    <linearGradient id="badgeGradient" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0%" stop-color="{{.Colors.Accent}}" stop-opacity="0.95"/>
      <stop offset="100%" stop-color="{{.Colors.Accent}}" stop-opacity="0.75"/>
    </linearGradient>
  </defs>

  <!-- Background -->
  <rect class="card" x="0.5" y="0.5" width="279" height="31"/>

  <!-- Text -->
  <text x="140" y="21" text-anchor="middle">
    OSS · {{.CompactText}}
  </text>
</svg>`

// detailedTemplate is the SVG template for the Detailed badge style (400x320)
const DefaultDetailed = `
<svg
  width="900"
  height="{{add 278 (mul (div (add (len .TopContributions) 2) 3) 120)}}"
  viewBox="0 0 900 {{add 278 (mul (div (add (len .TopContributions) 2) 3) 120)}}"
  xmlns="http://www.w3.org/2000/svg"
  role="img"
  aria-label="GitHub Open Source Contribution Stats">

  <!-- ========================= -->
  <!-- Theme + Styling -->
  <!-- ========================= -->
  <style>
    :root {
      --bg: {{.Colors.Background}};
      --bg-alt: {{.Colors.BackgroundAlt}};
      --text: {{.Colors.Text}};
      --text-muted: {{.Colors.TextSecondary}};
      --border: {{.Colors.Border}};
      --accent: {{.Colors.Accent}};
      --star: {{.Colors.Star}};
    }

    text {
      font-family: system-ui, -apple-system, BlinkMacSystemFont,
                   "Segoe UI", Helvetica, Arial, sans-serif;
    }

    .title {
      font-size: 26px;
      font-weight: 800;
      fill: var(--text);
      letter-spacing: -0.4px;
    }

    .subtitle {
      font-size: 13px;
      fill: var(--text-muted);
    }

    .card {
      fill: url(#cardGradient);
      stroke: var(--border);
      stroke-width: 1;
      rx: 14;
    }

    .card-glass {
      fill: url(#glass);
      rx: 14;
    }

    .metric-label {
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.1em;
      fill: var(--text-muted);
    }

    .metric-value {
      font-size: 28px;
      font-weight: 800;
      fill: var(--text);
    }

    .repo-name {
      font-size: 16px;
      font-weight: 700;
      fill: var(--text);
    }

    .repo-meta {
      font-size: 12px;
      fill: var(--text-muted);
    }

    .fade-in {
      opacity: 0;
      animation: fadeUp 0.6s ease-out forwards;
    }

    @keyframes fadeUp {
      from { opacity: 0; transform: translateY(6px); }
      to   { opacity: 1; transform: translateY(0); }
    }
  </style>

  <!-- ========================= -->
  <!-- Definitions -->
  <!-- ========================= -->
  <defs>
    <linearGradient id="cardGradient" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0%" stop-color="var(--bg-alt)" stop-opacity="0.8" />
      <stop offset="100%" stop-color="var(--bg-alt)" stop-opacity="0.6" />
    </linearGradient>

    <linearGradient id="glass" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0%" stop-color="var(--bg-alt)" stop-opacity="0.2" />
      <stop offset="100%" stop-color="var(--bg-alt)" stop-opacity="0" />
    </linearGradient>
  </defs>

  <!-- ========================= -->
  <!-- Background -->
  <!-- ========================= -->
  <rect
    x="0"
    y="0"
    width="900"
    height="{{add 278 (mul (div (add (len .TopContributions) 2) 3) 120)}}"
    rx="18"
    fill="var(--bg)"
  />

  <!-- ========================= -->
  <!-- Header -->
  <!-- ========================= -->
  <g class="fade-in">
    <text x="32" y="50" class="title">
      {{.Stats.Username}} · Open Source
    </text>
  </g>

  <!-- ========================= -->
  <!-- Metrics Row -->
  <!-- ========================= -->

  <!-- Projects -->
  <g class="fade-in" style="animation-delay: 100ms">
    <rect x="32" y="96" width="260" height="96" class="card"/>
    <rect x="32" y="96" width="260" height="96" class="card-glass"/>
    <text x="48" y="132" class="metric-label">Projects</text>
    <text x="48" y="168" class="metric-value">{{.TotalProjects}}</text>
  </g>

  <!-- PRs -->
  <g class="fade-in" style="animation-delay: 150ms">
    <rect x="320" y="96" width="260" height="96" class="card"/>
    <rect x="320" y="96" width="260" height="96" class="card-glass"/>
    <text x="336" y="132" class="metric-label">PRs Merged</text>
    <text x="336" y="168" class="metric-value">{{.TotalPRs}}</text>
  </g>

  <!-- Lines Changed -->
  <g class="fade-in" style="animation-delay: 200ms">
    <rect x="608" y="96" width="260" height="96" class="card"/>
    <rect x="608" y="96" width="260" height="96" class="card-glass"/>
    <text x="624" y="132" class="metric-label">Lines Changed</text>
    <text x="624" y="168" class="metric-value">{{.TotalLines}}</text>
  </g>

  <!-- ========================= -->
  <!-- Top Contributions -->
  <!-- ========================= -->
  <text x="32" y="224" class="subtitle fade-in" style="animation-delay: 260ms">
    Top Contributions
  </text>

  {{range $i, $r := .TopContributions}}
  {{ $col := mod $i 3 }}
  {{ $row := div $i 3 }}
  <g class="fade-in" style="animation-delay: {{add 300 (mul $i 80)}}ms">
    <rect
      x="{{add 32 (mul $col 288)}}"
      y="{{add 240 (mul $row 120)}}"
      width="260"
      height="108"
      class="card"
    />
    <rect
      x="{{add 32 (mul $col 288)}}"
      y="{{add 240 (mul $row 120)}}"
      width="260"
      height="108"
      class="card-glass"
    />

    <text
      x="{{add 48 (mul $col 288)}}"
      y="{{add 278 (mul $row 120)}}"
      class="repo-name">
      {{$r.RepoName}}
    </text>
    <text
      x="{{add 48 (mul $col 288)}}"
      y="{{add 304 (mul $row 120)}}"
      class="repo-meta">
      ⭐ {{$r.Stars}} · {{$r.PRs}} PRs Merged
    </text>
  </g>
  {{end}}

</svg>

`
