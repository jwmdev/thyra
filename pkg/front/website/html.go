// GENERATED BY textFileToGoConst
// GitHub:     github.com/logrusorgru/textFileToGoConst
// input file: html/front/website.html
// generated:  Mon Aug  1 10:10:59 CEST 2022

package website

const HTML = `<!DOCTYPE html>
<html>
	<head>
		<link
			href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css"
			rel="stylesheet"
			integrity="sha384-0evHe/X+R7YkIZDRvuzKMRqM+OrBnVFBL6DOitfPri4tjfHxaWutUpFmBp4vmVor"
			crossorigin="anonymous"
		/>
		<link rel="stylesheet" href="./website.css" />
	</head>
	<style></style>
	<body>
		<nav class="navbar navbar-expand-lg navbar-dark">
			<div class="container">
			  <a class="navbar-brand" href="#"><img src="./logo_massa.webp" class="massa-logo" alt="Massa logo" /></a>
			  <h2>Thyra</h1>
			  <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			  </button>

			  <div class="collapse navbar-collapse justify-content-end " id="navbarNav">
				<ul class="navbar-nav ">
				  <li class="nav-item">
					<a class="nav-link "  href="/wallet.mythyra.massa/index.html">Wallet</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link active" aria-current="page" href="#">Websites</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link" href="#">Pricing</a>
				  </li>
				</ul>
			  </div>
			</div>
		</nav>

		<div class="alert alert-danger" role="alert">
			This is a danger alert check it out!
		</div>

		<div class="container">

			<h2 class="mb-4 mt-5">Decentralized website storing</h1>

			<div class='row justify-content-md-center mt-5'>
				<div class='col-6'>
					<div id="create-form ">
						<div class="form-floating mb-3">
							<input
								class="form-control"
								placeholder="Website Name"
								id="websiteName"
								name="websiteName"
								type="text"
							/>
							<label for="websiteName">Website Name</label>
						</div>

						<div class="form-floating mb-4">
							<input
								class="form-control"
								placeholder="Website Description"
								id="websiteDescription"
								name="websiteDescription"
								type="text"
							/>
							<label for="websiteDescription">Website Description</label>
						</div>
					</div>
				</div>
			
				
			</div>

			<div class='row justify-content-md-center mt-5'>
				<div class='col-6 position-relative'>
					<input id='fileid' type='file' hidden/>
					<button id="updload-website" class="primary-button position-absolute top-50 start-50 translate-middle" id='buttonid' type='button' value='Upload MB' >Upload</button>
				</div>
			</div>
		</div>
	</body>
	<script
		src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js"
		integrity="sha384-pprn3073KE6tl6bjs2QrFaJGz5/SUsLqktiwsUTF55Jfv3qYSDhgCecCxMW52nD2"
		crossorigin="anonymous"
	></script>
	<script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
	<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
	<script src="website.js"></script>
</html>
`
