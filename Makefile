ifneq (,$(wildcard ./.env))
    include .env
    export
endif

clean:
	rm -f ${DATABASE_URL}
run:
	# TODO
	go build

# TODO run-clean
