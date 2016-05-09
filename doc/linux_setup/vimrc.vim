call plug#begin('~/.vim/plugged')
Plug 'Raimondi/delimitMate'
Plug 'Shougo/neocomplete'
Plug 'Xuyuanp/nerdtree-git-plugin'
Plug 'airblade/vim-gitgutter'
Plug 'bling/vim-airline'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'fatih/vim-go'
Plug 'scrooloose/nerdtree', { 'on':  'NERDTreeToggle' }
Plug 'scrooloose/syntastic'
Plug 'terryma/vim-multiple-cursors'
Plug 'tpope/vim-commentary'
Plug 'majutsushi/tagbar'
Plug 'nsf/gocode', {'rtp': 'vim'}
Plug 'ntpeters/vim-better-whitespace'
call plug#end()

" :PlugInstall
" :PlugClean
" :PlugUpdate
" :PlugUpgrade

" prepend in comparison with the line above
vnoremap . :norm.<CR>

" neocomplete
let g:acp_enableAtStartup = 0
let g:neocomplete#enable_at_startup = 1
let g:neocomplete#enable_smart_case = 1

" <TAB>: completion.
inoremap <expr><TAB>  pumvisible() ? "\<C-n>" :
	\ <SID>check_back_space() ? "\<TAB>" :
	\ neocomplete#start_manual_complete()
function! s:check_back_space() "{{{
let col = col('.') - 1
return !col || getline('.')[col - 1]  =~ '\s'
endfunction"}}}

" <space> + w to save the file
let mapleader = " "
nnoremap <Leader>w :w<CR>
nnoremap <Leader>q :q<CR>
nnoremap <Leader>e :SyntasticCheck gcc<CR>
nmap <Leader><Leader> V

" better line moves
map j gj
map k gk

map <Leader>a :bprev<Return>  " move buffer
map <Leader>s :bnext<Return>
map <Leader>d :bd<Return>     " close current buffer

" vim-Go
filetype plugin on
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1
let g:go_highlight_interfaces = 1
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1
let g:go_fmt_command = "goimports"
let g:go_fmt_autosave = 1

let g:gitgutter_max_signs = 100000

" https://github.com/scrooloose/syntastic
set laststatus=2
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*

let g:syntastic_always_populate_loc_list = 1
let g:syntastic_auto_loc_list = 1
let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 1
let g:syntastic_mode_map = {
    \ "mode": "active",
    \ "passive_filetypes": ["go", "asm"] }

let g:syntastic_cpp_compiler_options = '-std=c++11'
let g:syntastic_go_checkers = ['golint', 'govet', 'errcheck']

let g:airline_detect_modified=1
let g:airline_detect_paste=1
let g:airline_section_b = '%{strftime("%c")}'
let g:airline#extensions#tabline#enabled = 1
let g:airline_powerline_fonts=0

if !exists('g:airline_symbols')
  let g:airline_symbols = {}
endif
let g:airline_symbols.space = "\ua0"

" NERDTree
" map <F2> :NERDTreeToggle<cr>
" map <F3> <C-w><C-w>
"
map <ESC>2 :NERDTreeToggle<cr>
map <ESC>3 <C-w><C-w>
map <ESC>4 :TagbarToggle<CR>

"http://nvie.com/posts/how-i-boosted-my-vim/
" set wrap
" set tw=79
set formatoptions+=t
set tabstop=4     " a tab is four spaces
set backspace=indent,eol,start
                  " allow backspacing over everything in insert mode
set autoindent    " always set autoindenting on
set copyindent    " copy the previous indentation on autoindenting
set number        " always show line numbers
set shiftwidth=4  " number of spaces to use for autoindenting
set shiftround    " use multiple of shiftwidth when indenting with '<' and '>'
set showmatch     " set show matching parenthesis
set ignorecase    " ignore case when searching
set smartcase     " ignore case if search pattern is all lowercase,
                  " case-sensitive otherwise
set smarttab      " insert tabs on the start of a line according to
                  " shiftwidth, not tabstop

set incsearch     " show search matches as you type
set hlsearch      " highlight all search matches
" Use <C-L> to clear the highlighting of :set hlsearch.
if maparg('<C-L>', 'n') ==# ''
  nnoremap <silent> <C-L> :nohlsearch<C-R>=has('diff')?'<Bar>diffupdate':''<CR><CR><C-L>
endif

set autoread

set noswapfile

if &encoding ==# 'latin1' && has('gui_running')
  set encoding=utf-8
endif

" In many terminal emulators the mouse works just fine, thus enable it.
if has('mouse')
  set mouse=a
endif

set nolazyredraw " don't redraw while executing macros

" https://github.com/philips/etc
autocmd BufNewFile,BufRead *.md,*.markdown,*.mdown,*.mkd,*.mkdn set filetype=markdown
au FileType python setl autoindent tabstop=4 expandtab shiftwidth=4 softtabstop=4
au FileType html setl autoindent tabstop=2 expandtab shiftwidth=2 softtabstop=2

autocmd FileType c,cpp,go,java setlocal commentstring=//\ %s
let g:delimitMate_expand_cr=1

" spellcheck default for markdown files
autocmd BufRead,BufNewFile *.md setlocal spell

" F5 to toggle spell check
map <F5> :setlocal spell! spelllang=en_us<CR>

" syntax off
" colorscheme darkblue
syntax enable
set background=dark
colorscheme default

function! Multiple_cursors_before()
    exe 'NeoCompleteLock'
    echo 'Disabled autocomplete'
endfunction

function! Multiple_cursors_after()
    exe 'NeoCompleteUnlock'
    echo 'Enabled autocomplete'
endfunction

" Allow saving of files as sudo when I forgot to start vim using sudo.
cmap w!! w !sudo tee > /dev/null %

let g:tagbar_type_go = {
    \ 'ctagstype' : 'go',
    \ 'kinds'     : [
        \ 'p:package',
        \ 'i:imports:1',
        \ 'c:constants',
        \ 'v:variables',
        \ 't:types',
        \ 'n:interfaces',
        \ 'w:fields',
        \ 'e:embedded',
        \ 'm:methods',
        \ 'r:constructor',
        \ 'f:functions'
    \ ],
    \ 'sro' : '.',
    \ 'kind2scope' : {
        \ 't' : 'ctype',
        \ 'n' : 'ntype'
    \ },
    \ 'scope2kind' : {
        \ 'ctype' : 't',
        \ 'ntype' : 'n'
    \ },
    \ 'ctagsbin'  : 'gotags',
    \ 'ctagsargs' : '-sort -silent'
\ }

