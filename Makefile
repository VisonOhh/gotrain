# projects
cobra_tour := projects/cobra_tour

projects := $(cobra_tour)

define make_seq
	for subdir in $1 ; do				\
		cd $$subdir && make $2;			\
		ret_code=$$?;					\
		if [ $$ret_code != 0 ]; then	\
			exit $$ret_code;			\
		fi;								\
		cd -;							\
	done;
endef

.PHONY: all $(projects)
all: build

build:
	$(call make_seq , $(projects) , build)
test:
	$(call make_seq , $(projects) , test)
clean:
	@for d in $(projects);        		\
	do                                  \
		$(MAKE) --directory=$$d clean;  \
	done;
