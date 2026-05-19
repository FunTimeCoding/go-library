package conversations

const inlineCSS = `
	html, body {
		height: 100%;
		margin: 0;
		overflow: hidden;
	}
	.conversation-layout {
		display: grid;
		grid-template-columns: 350px 1fr;
		height: 100vh;
	}
	.sidebar {
		overflow-y: auto;
		border-right: 1px solid var(--pico-muted-border-color);
		padding: 0.5rem;
	}
	.sidebar-entry {
		padding: 0.5rem;
		cursor: pointer;
		border-radius: var(--pico-border-radius);
		margin-bottom: 0.25rem;
	}
	.sidebar-entry:hover {
		background: var(--pico-muted-border-color);
	}
	.sidebar-entry small {
		color: var(--pico-muted-color);
		display: block;
	}
	.entry-name {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.rename-icon {
		opacity: 0;
		cursor: pointer;
		font-size: 0.85rem;
	}
	.sidebar-entry:hover .rename-icon {
		opacity: 0.5;
	}
	.rename-icon:hover {
		opacity: 1 !important;
	}
	.rename-input {
		width: 100%;
		font-size: 0.9rem;
		padding: 0.2rem 0.4rem;
		margin: 0;
	}
	.panel {
		overflow-y: auto;
		padding: 1rem 2rem;
	}
	.panel-placeholder {
		color: var(--pico-muted-color);
		text-align: center;
		margin-top: 40vh;
	}
	.message {
		margin-bottom: 1rem;
		padding: 0.75rem 1rem;
		border-radius: var(--pico-border-radius);
		max-width: 1000px;
		margin-left: auto;
		margin-right: auto;
	}
	.message-user {
		border-left: 3px solid #0074d9;
	}
	.message-assistant {
		border-left: 3px solid #2ecc40;
		background: var(--pico-card-background-color);
	}
	.message-role {
		font-size: 0.75rem;
		font-weight: bold;
		text-transform: uppercase;
		color: var(--pico-muted-color);
		margin-bottom: 0.25rem;
	}
	.message-text {
		white-space: pre-wrap;
		word-wrap: break-word;
	}
	.sidebar-filter {
		width: 100%;
		padding: 0.4rem 0.5rem;
		margin-bottom: 0.5rem;
		font-size: 0.85rem;
	}
`

const filterJS = `
	function filterSidebar(query) {
		var entries = document.querySelectorAll('.sidebar-entry');
		var lower = query.toLowerCase();
		entries.forEach(function(e) {
			var name = e.querySelector('.entry-name span');
			if (!name) return;
			e.style.display = name.textContent.toLowerCase().indexOf(lower) !== -1 ? '' : 'none';
		});
	}
`

const infiniteScrollJS = `
	(function() {
		var sidebar = document.querySelector('.sidebar');
		if (!sidebar) return;
		var loading = false;
		sidebar.addEventListener('scroll', function() {
			if (loading) return;
			if (sidebar.scrollTop + sidebar.clientHeight < sidebar.scrollHeight - 50) return;
			var sentinel = sidebar.querySelector('[data-load-more]');
			if (!sentinel) return;
			loading = true;
			var url = sentinel.getAttribute('data-load-more');
			fetch(url).then(function(r) { return r.text(); }).then(function(html) {
				sentinel.outerHTML = html;
				loading = false;
			});
		});
	})();
`

const scrollToBottomJS = `
	document.addEventListener('htmx:afterSwap', function(e) {
		if (e.detail.target.id === 'panel') {
			e.detail.target.scrollTop = e.detail.target.scrollHeight;
		}
	});
`
