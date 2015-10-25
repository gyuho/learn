" Start of Vundle.vim
set nocompatible
filetype off

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" let Vundle manage Vundle, required
Plugin 'gmarik/Vundle.vim'

" Plugins
Plugin 'airblade/vim-gitgutter'
Plugin 'bling/vim-airline'
Plugin 'ctrlpvim/ctrlp.vim'
Plugin 'fatih/vim-go'
Plugin 'godlygeek/tabular'
Plugin 'Raimondi/delimitMate'
Plugin 'scrooloose/nerdtree'
Plugin 'scrooloose/syntastic'
Plugin 'majutsushi/tagbar'
Plugin 'terryma/vim-multiple-cursors'
Plugin 'tpope/vim-commentary'
Plugin 'Valloric/YouCompleteMe'

call vundle#end()
filetype plugin indent on
" End of Vundle.vim

syntax on

" to prepend in comparison with the line above
vnoremap . :norm.<CR>

" http://sheerun.net/2014/03/21/how-to-boost-your-vim-productivity/
let mapleader = " "

" <Space> + w to save the file
nnoremap <Leader>w :w<CR>
nnoremap <Leader>q :q<CR>
nnoremap <Leader>e :SyntasticCheck gcc<CR>
nmap <Leader><Leader> V

" Move line by line even when the line is wrapped
map j gj
map k gk

" Move between buffers
map <Leader>a :bprev<Return>
map <Leader>s :bnext<Return>
" Close the current buffer
map <Leader>d :bd<Return>

" tabular
nnoremap <Leader>t= :Tabularize /=<CR>
nnoremap <Leader>t: :Tabularize /:\zs<CR>
" :Tab/|
" for |

" Vim-Go
let g:go_fmt_command = "goimports"
let g:go_fmt_autosave = 1
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1
let g:go_doc_keywordprg_enabled = 0
let g:go_def_mapping_enabled = 0

" https://github.com/scrooloose/syntastic
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*

let g:syntastic_always_populate_loc_list = 1
let g:syntastic_auto_loc_list = 1
let g:syntastic_check_on_open = 1
let g:syntastic_mode_map = {
    \ "mode": "active",
    \ "passive_filetypes": ["go", "asm"] }

let g:syntastic_cpp_compiler_options = '-std=c++11'
let g:airline#extensions#tabline#enabled = 1
set laststatus=2

" NERDTree
map <F2> :NERDTreeToggle<cr>
map <F3> <C-w><C-w>

"http://nvie.com/posts/how-i-boosted-my-vim/
set wrap
set tw=79
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

set noswapfile

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

" tagbar
nmap <F8> :TagbarToggle<CR>

set t_Co=256
colorscheme default

