APP_NAME := ogm-tag
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

.PHONY: build
build: 
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run: 
	./bin/${APP_NAME}

.PHONY: install
install: 
	go install

.PHONY: clean
clean: 
	rm -rf /tmp/ogm-tag.db

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"T_01", "name":"处理器", "flag": 1024, "alias":"CPU"}'
	# 添加不存在的标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"T_01", "name":"处理器", "flag": 1024, "alias":"CPU"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"T_02", "name":"硬盘", "flag": 1024, "alias":"HD"}'
	# 添加已存在的标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"T_01", "name":"处理器", "flag": 1024, "alias":"CPU"}'
	# 列举
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.ListTag ''
	# 更新不存在的标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.UpdateTag '{"code":"T_09", "name":"移动处理器", "flag": 2048, "alias":{"en_US":"Mobile CPU", "zh_CN":"移动处理器"}}'
	# 更新存在的标签的部分字段
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.UpdateTag '{"code":"T_01", "name":"移动处理器", "flag": 2048 }'
	# 更新存在的标签的全部字段
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.UpdateTag '{"code":"T_01", "name":"骁龙835", "flag": 2048, "alias":{"en_US":"Snapdragon 835", "zh_CN":"骁龙 835"}, "keyword":["cpu","处理器","移动端","mobile"]}'
	# 列举
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.ListTag ''
	# 智能提示编号
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SuggestFilter '{"input":"01"}'
	# 智能提示名称
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SuggestFilter '{"input":"移动"}'
	# 智能提示关键字
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SuggestFilter '{"input":"MOB"}'
	# 智能提示别名
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SuggestFilter '{"input":"dra"}'
	# 搜索编号 
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SearchTag '{"filter":"01"}'
	# 搜索名称
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SearchTag '{"filter":"骁龙835"}'
	# 搜索别名
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SearchTag '{"filter":"Snapdragon 835"}'
	# 搜索关键字
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.SearchTag '{"filter":"mobile"}'
	# 添加受体, 标签不存在
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"T_09", "owner":"001"}'
	# 添加受体
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"T_01", "owner":"001"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"T_01", "owner":"002"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"T_02", "owner":"001"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"T_02", "owner":"003"}'
	# 检索受体
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.FilterTag '{"code":["T_09"]}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.FilterTag '{"code":["T_01"]}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.FilterTag '{"code":["T_02"]}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.FilterTag '{"code":["T_01", "T_02"]}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.FilterTag '{"code":["T_01", "T_09"]}'
	# 列举受体的标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.ListTag '{"owner":"001"}'

	# 删除不存在的标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.RemoveTag '{"code":"T_09"}'
	# 删除存在的标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.RemoveTag '{"code":"T_01"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.RemoveTag '{"code":"T_02"}'

.PHONY: tcall
tcall:
	mkdir -p ./bin
	go build -o ./bin/ ./tester
	./bin/tester

.PHONY: fill
fill:
	# 添加标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"A_01", "name":"Snapdragon 835", "flag": 1024, "alias":"骁龙835"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"A_02", "name":"Intel i5", "flag": 1024, "alias":"英特尔i5"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.AddTag '{"code":"B_01", "name":"SSD 128G", "flag": 1024, "alias":"128G固态"}'
	# 更新标签
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.UpdateTag '{"code":"A_01", "keyword":["cpu","处理器","移动端","mobile"]}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Collection.UpdateTag '{"code":"A_02", "keyword":["cpu","处理器","桌面端","PC"]}'
	# 添加受体
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"A_01", "owner":"HUAWEI-P30"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"A_02", "owner":"Dell-XPS13"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"B_01", "owner":"HUAWEI-P30"}'
	MICRO_REGISTRY=consul micro call xtc.api.ogm.tag Dummy.AddTag '{"code":"B_01", "owner":"Dell-XPS13"}'

.PHONY: post
post:
	curl -X POST -d '{"msg":"hello"}' 127.0.0.1:8080/ogm/tag/Healthy/Echo

.PHONY: bm
bm:
	python3 ./benchmark.py

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

