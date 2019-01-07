build_js:
	gopherjs build -o build/js/adventure-island.js
	echo '<!DOCTYPE html><script src="adventure-island.js"></script>' > build/js/index.html

serve_js:
	echo "https://localhost:8000"
	python3 -m http.server --directory build/js/


clean:
	rm -rf ./build/