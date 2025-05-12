String getTenantFromSubdomain(Uri uri) {
  final host = uri.host;
  final parts = host.split('.');
  return parts.length > 2 ? parts[0] : 'default';
}
