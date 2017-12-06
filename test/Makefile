# Makefile for tests/

PYBIN		?= python
MODULES		= common directfs director filesystem rdbplugin
MODDIR		:= $(shell cd $(CURDIR)/..; pwd)

## The PYTHONPATH is complicated by the directory layout.
# We need both the full path (common/v1/common_pb2.py) and the short
# path (./common_pb2.py) so we can easily find the modules from
# Python programs.
# The generated gRPC/Protobuf code uses the short form.

# Hack to get a colon-separated list of paths.
EMPTY		:=
SPACE		:= $(EMPTY) $(EMPTY)
MOD_PP		:= $(subst $(SPACE),:,$(foreach mod,$(MODULES),$(MODDIR)/$(mod)/v1))

PYTHONPATH	:= $(MODDIR):$(MOD_PP):$(CURDIR)

TESTS		:= $(wildcard test_*.py)
TESTOPTS	?= -v

PYENV		:= PYTHONPATH=$(PYTHONPATH)
PYRUN		:= env $(PYENV) $(PYBIN)


all: test

test: $(TESTS)
	$(PYRUN) -m unittest discover $(CURDIR) $(TESTOPTS)
